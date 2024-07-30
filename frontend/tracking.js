(function () {
  const backendUrl = "http://localhost:8080/track";
  let startTime = new Date().getTime();

  const sendData = (data) => {
    console.log("event_type: ", data.type);
    fetch(backendUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).catch((error) => console.error("Error sending data:", error));
  };

  const getUserId = () => {
    let userId = localStorage.getItem("analytics_user_id");
    if (!userId) {
      userId = "_" + Math.random().toString(36).substr(2, 9);
      localStorage.setItem("analytics_user_id", userId);
    }
    return userId;
  };

  const trackPageView = () => {
    const data = {
      type: "pageview",
      url: window.location.href,
      referrer: document.referrer,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackClick = (event) => {
    const data = {
      type: "click",
      x: event.clientX,
      y: event.clientY,
      target: event.target.tagName,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackDuration = () => {
    const endTime = new Date().getTime();
    const duration = Math.round((endTime - startTime) / 1000); // Duration in seconds
    const data = {
      type: "duration",
      url: window.location.href,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
      duration: duration,
    };
    sendData(data);
  };

  window.addEventListener("load", trackPageView);
  window.addEventListener("click", trackClick);
  window.addEventListener("beforeunload", trackDuration);
})();
