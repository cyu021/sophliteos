
const wsUrl = "ws://${window.location.host}/ws"

let websocket = null;// new WebSocket(wsUrl);

const create = (url) => {
    websocket = new WebSocket(url);
}

const connect = () => {
    return new Promise((resolve, reject) => {
        
        console.log('websocket.before state:', websocket.readyState, websocket.url);

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

const close = () => {
    websocket.close(1000, 'mannualy close');
}

const ws = {
    create,
    connect,
    listen,
    close,
}
  
export default ws;