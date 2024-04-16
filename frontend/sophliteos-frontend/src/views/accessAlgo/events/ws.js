
const wsUrl = "ws://${window.location.host}/ws"

let websocket = null;// new WebSocket(wsUrl);

const create = (url) => {
    websocket = new WebSocket(url.replace(":8080", ":5566"));
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

const ws = {
    create,
    connect,
    listen,
}
  
export default ws;