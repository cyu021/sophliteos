
const wsUrl = "ws://192.168.0.188:5566/ws"

const websocket = new WebSocket(wsUrl);

const connect = () => {
    return new Promise((resolve, reject) => {
        
        console.log('websocket.onopen', websocket.readyState);

        if(websocket.readyState === 1) {
            resolve();
            return;
        }

        websocket.onopen = () => {
            console.log("websocket connected");
            resolve();
        }

        websocket.onclose = () => {
            console.log("websocket closed");
            
        }
    })
}

const listen = (callback) => {
    console.log('websocket.onmessage');

    websocket.onmessage = (event) => {
        console.log('收到消息:', event.data);  
        callback(JSON.parse(event.data));
    }
}

const ws = {
    connect,
    listen,
}
  
export default ws;