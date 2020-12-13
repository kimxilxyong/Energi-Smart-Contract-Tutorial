/* SPDX-License-Identifier: MIT
   Copyright (c) 2020 Kim Il Yong */


export const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text).then(
        () => {
            console.log("Copied", text, "to clipboard");
        },
        (e) => {
            console.log("Copy to clipboard failed", e);
        },

    );
};

export const getTime = () => {

  let currentTime = new Date();
  let hours = currentTime.getHours();
  let minutes = currentTime.getMinutes();
  let seconds = currentTime.getSeconds();

  if (minutes < 10) {
    minutes = "0" + minutes;
  }
  if (seconds < 10) {
    seconds = "0" + seconds;
  }
  return hours + ":" + minutes + ":" + seconds;
};

export const timeDifference = (current, previous) => {

  const msPerMinute = 60 * 1000;
  const msPerHour = msPerMinute * 60;
  const msPerDay = msPerHour * 24;
  const msPerMonth = msPerDay * 30;
  const msPerYear = msPerDay * 365;

  const elapsed = current - previous;

  if (elapsed < msPerMinute) {
    return Math.round(elapsed / 1000) + ' seconds ago';
  } else if (elapsed < msPerHour) {
    return Math.round(elapsed / msPerMinute) + ' minutes ago';
  } else if (elapsed < msPerDay) {
    return Math.round(elapsed / msPerHour) + ' hours ago';
  } else if (elapsed < msPerMonth) {
    return 'approximately ' + Math.round(elapsed / msPerDay) + ' days ago';
  } else if (elapsed < msPerYear) {
    return 'approximately ' + Math.round(elapsed / msPerMonth) + ' months ago';
  } else {
    return 'approximately ' + Math.round(elapsed / msPerYear) + ' years ago';
  }
};

export const scrollTo = (id) => {
  //get the element to be scrolled
  let sct = document.getElementById(id);
  //make the parent to scroll into view, smoothly!
  sct.scrollIntoView({ block: "end", behavior: "smooth", });
  console.log("Scroll ", id, sct);
};
