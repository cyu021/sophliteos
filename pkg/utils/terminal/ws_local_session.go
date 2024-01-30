package terminal

import (
	"encoding/base64"
	"encoding/json"
	"sophliteos/logger"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type LocalWsSession struct {
	slave  *LocalCommand
	wsConn *websocket.Conn

	writeMutex sync.Mutex
}

func NewLocalWsSession(cols, rows int, wsConn *websocket.Conn, slave *LocalCommand) (*LocalWsSession, error) {
	if err := slave.ResizeTerminal(cols, rows); err != nil {
		logger.Error("ssh pty change windows size failed, err: %v", err)
	}

	return &LocalWsSession{
		slave:  slave,
		wsConn: wsConn,
	}, nil
}

func (sws *LocalWsSession) Start(quitChan chan bool) {
	go sws.handleSlaveEvent(quitChan)
	go sws.receiveWsMsg(quitChan)
}

// 从终端读取数据，返回页面
func (sws *LocalWsSession) handleSlaveEvent(exitCh chan bool) {
	defer setQuit(exitCh)
	defer logger.Info("发送ws信息线程已退出")

	buffer := make([]byte, 1024)
	for {
		select {
		case <-exitCh:
			return
		default:
			n, err := sws.slave.Read(buffer)
			if err != nil {
				logger.Error("read from docker err : %s", err.Error())
				return
			}
			_ = sws.masterWrite(buffer[:n])

		}
	}
}

func (sws *LocalWsSession) masterWrite(data []byte) error {
	sws.writeMutex.Lock()
	defer sws.writeMutex.Unlock()
	wsData, err := json.Marshal(WsMsg{
		Type: WsMsgCmd,
		Data: base64.StdEncoding.EncodeToString(data),
	})
	if err != nil {
		return errors.Wrapf(err, "failed to encoding to json")
	}
	err = sws.wsConn.WriteMessage(websocket.TextMessage, wsData)
	if err != nil {
		return errors.Wrapf(err, "failed to write to master")
	}
	return nil
}

// 接收页面消息,转发到终端
func (sws *LocalWsSession) receiveWsMsg(exitCh chan bool) {
	wsConn := sws.wsConn
	defer setQuit(exitCh)
	defer logger.Info("接收ws信息线程已退出")
	for {
		select {
		case <-exitCh:
			return
		default:
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				logger.Error("reading webSocket message failed, err: %v", err)
				return
			}
			msgObj := WsMsg{}
			_ = json.Unmarshal(wsData, &msgObj)
			switch msgObj.Type {
			case WsMsgResize:
				if msgObj.Cols > 0 && msgObj.Rows > 0 {
					if err := sws.slave.ResizeTerminal(msgObj.Cols, msgObj.Rows); err != nil {
						logger.Error("ssh pty change windows size failed, err: %v", err)
					}
				}
			case WsMsgCmd:
				decodeBytes, err := base64.StdEncoding.DecodeString(msgObj.Data)
				if err != nil {
					logger.Error("websock cmd string base64 decoding failed, err: %v", err)
				}
				sws.sendWebsocketInputCommandToSshSessionStdinPipe(decodeBytes)
			case WsMsgHeartbeat:
				// 接收到心跳包后将心跳包原样返回，可以用于网络延迟检测等情况
				err = wsConn.WriteMessage(websocket.TextMessage, wsData)
				if err != nil {
					logger.Error("ssh sending heartbeat to webSocket failed, err: %v", err)
				}
			}
		}
	}
}

func (sws *LocalWsSession) sendWebsocketInputCommandToSshSessionStdinPipe(cmdBytes []byte) {
	if _, err := sws.slave.Write(cmdBytes); err != nil {
		logger.Error("ws cmd bytes write to ssh.stdin pipe failed, err: %v", err)
	}
}
