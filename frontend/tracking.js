(function () {
  const BASE_URL = "http://localhost:8080/track";

  function sendData(data) {
    fetch(BASE_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).catch((error) => console.error("Error sending data:", error));
  }

  // Track page views
  sendData({
    eventType: "page_view",
    url: window.location.href,
    timestamp: new Date().toISOString(),
  });

  // Track clicks
  document.addEventListener("click", function (event) {
    if (event.target.classList.contains("trackable")) {
      sendData({
        eventType: "click",
        element: event.target.tagName,
        elementId: event.target.id,
        timestamp: new Date().toISOString(),
      });
    }
  });

  // Track session duration
  let startTime = new Date().getTime();
  window.addEventListener("beforeunload", function () {
    let endTime = new Date().getTime();
    let duration = endTime - startTime;
    sendData({
      eventType: "session_duration",
      duration: duration,
      timestamp: new Date().toISOString(),
    });
  });

  // Track custom events (e.g., button clicks, form submissions)
  function trackEvent(eventType, details) {
    sendData({
      eventType: eventType,
      details: details,
      timestamp: new Date().toISOString(),
    });
  }
})();
