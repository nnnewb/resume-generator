"use strict";

(function () {
  var ws = new WebSocket("ws://localhost:8889/ws");
  ws.onopen = function (ev) {
    setInterval(function () {
      ws.send(JSON.stringify({ alive: true }));
    }, 1000);
  };

  ws.onerror = function (err) {
    console.error(err);
    alert("websocket 错误, 已停止热重载, 请手动刷新页面");
  };

  ws.onmessage = function (msg) {
    try {
      var m = JSON.parse(msg.data);
      if (m.reload) {
        window.location.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };
})();
