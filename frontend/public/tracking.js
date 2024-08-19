(function () {
  const backendUrl = "http://localhost:8080/track";
  let startTime = new Date().getTime();
  let idleStartTime = null;
  const idleTimeout = 60000;
  let isFieldFocused = false;
  let currentFocusedElement = null;

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
    if (!isFieldFocused) {
      const data = {
        type: "click",
        x: event.clientX,
        y: event.clientY,
        target: event.target.tagName,
        userId: getUserId(),
        timestamp: new Date().toISOString(),
      };
      sendData(data);
    }
    isFieldFocused = false;
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

  const trackScroll = () => {
    const scrollTop = window.scrollY || document.documentElement.scrollTop;
    const windowHeight = window.innerHeight;
    const documentHeight = document.documentElement.scrollHeight;
    const scrollPercentage = Math.round(
      ((scrollTop + windowHeight) / documentHeight) * 100
    );
    const data = {
      type: "scroll",
      url: window.location.href,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
      scrollPercentage: scrollPercentage,
    };
    sendData(data);
  };

  const trackElementScroll = (element) => {
    const scrollTop = element.scrollTop;
    const elementHeight = element.scrollHeight - element.clientHeight;
    const scrollPercentage = Math.round((scrollTop / elementHeight) * 100);

    const data = {
      type: "element_scroll",
      url: window.location.href,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
      elementId: element.id || null,
      scrollPercentage: scrollPercentage,
    };
    sendData(data);
  };

  const trackMouseMovement = (event) => {
    const data = {
      type: "mousemove",
      x: event.clientX,
      y: event.clientY,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackHover = (event) => {
    const data = {
      type: "hover",
      target: event.target.tagName,
      id: event.target.id || null,
      className: event.target.className || null,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackFormSubmission = (event) => {
    const data = {
      type: "form_submission",
      formId: event.target.id || null,
      formClassName: event.target.className || null,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  // Monitor when users focus on or leave the form fields
  const trackFieldFocus = (event) => {
    if (currentFocusedElement !== event.target) {
      isFieldFocused = true;
      currentFocusedElement = event.target;
    }
    // isFieldFocused = true;
    const data = {
      type: "field_focus",
      fieldId: event.target.id || null,
      fieldName: event.target.name || null,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackFieldBlur = (event) => {
    setTimeout(() => {
      if (
        document.activeElement &&
        document.activeElement.form === event.target.form
      ) {
        return;
      }
      currentFocusedElement = null;
      const data = {
        type: "field_blur",
        fieldId: event.target.id || null,
        fieldName: event.target.name || null,
        userId: getUserId(),
        timestamp: new Date().toISOString(),
      };
      sendData(data);
    }, 0);
  };

  const trackIdleTime = (idleDuration) => {
    const data = {
      type: "idle_time",
      duration: Math.round(idleDuration / 1000), // Duration in seconds
      userId: getUserId(),
      timestamp: new Date().toISOString(),
      url: window.location.href,
    };
    sendData(data);
  };

  const startIdleTimer = () => {
    idleStartTime = new Date().getTime();
    setTimeout(checkIdleTime, idleTimeout);
  };

  const checkIdleTime = () => {
    const currentTime = new Date().getTime();
    if (currentTime - idleStartTime >= idleTimeout) {
      trackIdleTime(currentTime - idleStartTime);
    } else {
      startIdleTimer();
    }
  };

  const resetIdleTimer = () => {
    idleStartTime = new Date().getTime();
  };

  const trackVideoPlay = (event) => {
    const data = {
      type: "video_play",
      videoId: event.target.id || null,
      videoUrl: event.target.src,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  const trackVideoCompletion = (event) => {
    if (event.target.currentTime === event.target.duration) {
      const data = {
        type: "video_completion",
        videoId: event.target.id || null,
        videoUrl: event.target.src,
        userId: getUserId(),
        timestamp: new Date().toISOString(),
      };
      sendData(data);
    }
  };

  // Track audio plays
  const trackAudioPlay = (event) => {
    const data = {
      type: "audio_play",
      audioId: event.target.id || null,
      audioUrl: event.target.src,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  // Track downloads
  const trackDownload = (event) => {
    const data = {
      type: "download",
      downloadUrl: event.target.href,
      userId: getUserId(),
      timestamp: new Date().toISOString(),
    };
    sendData(data);
  };

  // Track image views (for lazy-loaded images)
  const trackImageView = (event) => {
    if (event.target.tagName.toLowerCase() === "img" && event.target.complete) {
      const data = {
        type: "image_view",
        imageUrl: event.target.src,
        userId: getUserId(),
        timestamp: new Date().toISOString(),
      };
      sendData(data);
    }
  };

  window.addEventListener("load", trackPageView);
  window.addEventListener("click", trackClick);
  window.addEventListener("beforeunload", trackDuration);
  window.addEventListener("scroll", trackScroll);

  // Track element scroll
  const scrollableElements = document.querySelectorAll("div, p, section");
  console.log("scrolling");
  scrollableElements.forEach((element) => {
    if (element.scrollHeight > element.clientHeight) {
      element.addEventListener("scroll", () => trackElementScroll(element));
    }
  });

  window.addEventListener("mousemove", resetIdleTimer);
  window.addEventListener("mousemove", trackMouseMovement);
  window.addEventListener("keypress", resetIdleTimer);

  window.addEventListener("mouseover", trackHover, true);
  document.addEventListener("submit", trackFormSubmission);
  document.addEventListener("focus", trackFieldFocus, true);
  document.addEventListener("blur", trackFieldBlur, true);

  // Media events
  document.addEventListener("play", trackVideoPlay, true);
  document.addEventListener("ended", trackVideoCompletion, true);
  document.addEventListener("play", trackAudioPlay, true);

  // Download tracking
  document.addEventListener("click", function (event) {
    if (
      (event.target.tagName.toLowerCase() === "a" && event.target.href) ||
      (event.target.tagName.toLowerCase() === "button" &&
        event.target.classList.contains("button"))
    ) {
      trackDownload(event);
    }
  });

  // Image view tracking
  document.addEventListener("load", trackImageView, true);

  startIdleTimer();
})();
