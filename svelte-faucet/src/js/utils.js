/* SPDX-License-Identifier: MIT
   Copyright (c) 2020 Kim Il Yong */


/* Get the country flag
   1be9a6884abd4c3ea143b59ca317c6b2
   645728ce29a740769339b7493f06fa17
   https://ipgeolocation.abstractapi.com/v1/?api_key=1be9a6884abd4c3ea143b59ca317c6b2&ip_address=127.5.1.2
   https://ipgeolocation.abstractapi.com/v1/?api_key=1be9a6884abd4c3ea143b59ca317c6b2&ip_address=2600:8802:901:1500:24f3:91fa:64cf:bec2
*/

export const getCountryFlag = async () => {
    try {
        let r = await fetch("https://ipgeolocation.abstractapi.com/v1/?api_key=645728ce29a740769339b7493f06fa17", {
            credentials: 'omit',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json'
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            referrerPolicy: 'no-referrer',
        });
        // TODO REMOVE debug
        console.log("Flag raw:", r);

        let flag = "🇦🇶";
        if (r) {
            try {
                r = r.json();
            } catch (e) {
                return "🏴󠁵󠁳󠁴󠁸󠁿";
            }
            if (r.error) {
                if (r.error.details) {
                    if (r.error.details.ip_address) {
                        flag = "🚣";
                    } else {
                        flag = "🏴‍☠️";
                    }
                } else {
                    flag = "🎌";
                }
            } else {
                try {
                    flag = r.flag.emoji;
                    if (r.security && r.security.is_vpn && r.security.is_vpn === true) {
                        flag = "🏴‍☠️" + flag;
                    }
                    // TODO REMOVE debug
                    console.log("Flag JSON", r.json());
                } catch (e) {
                    if (!flag) {
                        flag = "🏳️";
                    }
                }
            }
        } else {
            flag = "🇺🇳";
        }
        return flag;

    } catch (e) {
        //throw (e);
        //return "🎌";
        return "🏴󠁵󠁳󠁴󠁸󠁿";
    }

}

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

const stringInject = (str = '', obj = {}) => {
  let newStr = str;
  Object.keys(obj).forEach((key) => {
    let placeHolder = `#${key}#`;
    if (newStr.includes(placeHolder)) {
      newStr = newStr.replace(placeHolder, obj[key] || " ");
    }
  });
  return newStr;
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

  const msPerSecond = 1000;
  const msPerMinute = 60 * msPerSecond;
  const msPerHour = msPerMinute * 60;
  const msPerDay = msPerHour * 24;
  const msPerWeek = msPerDay * 7;
  const msPerMonth = msPerDay * 30;
  const msPerYear = msPerDay * 365;
  const msPerCentury = msPerYear * 100;
  const second = "just now";
  const minute = "#count# second#one# ago";
  const hour = "#count# minute#one# ago";
  const day = "#count# hour#one# ago";
  const week = "#count# day#one# ago";
  const month = "#count# week#one# ago";
  const year = "#count# month#one# ago";
  const century = "#count# year#one# ago";

  const units = [{ ms: msPerSecond, u: second, },
                { ms: msPerMinute, u: minute, },
                { ms: msPerHour, u: hour, },
                { ms: msPerDay, u: day, },
                { ms: msPerWeek, u: week, },
                { ms: msPerMonth, u: month, },
                { ms: msPerYear, u: year, },
                { ms: msPerCentury, u: century, },
               ];

  const elapsed = current - previous;
  //console.log("timeDifference: current", current, "previous", previous, "elapsed", elapsed);

  let result = "";
  let ago = 101;

  for (let i = 0; i < units.length; i++) {
    let isOne = "s";
    if (elapsed < units[i].ms) {
      if (i > 0) {
        ago = Math.round(elapsed / units[i - 1].ms);
      }
      if (ago === 1) {
        isOne = "";
      }
      result = stringInject(units[i].u, {count: ago, one: isOne,});
      //console.log("AGO Result:", ago, result);
      break;
    }
  }
  if (result === "") {
    result = "More than 100 years ago";
  }
  return result;

/*   if (elapsed < msPerMinute) {
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
  } */
};

export const scrollTo = (id) => {
  //get the element to be scrolled
  let sct = document.getElementById(id);
  //make the parent to scroll into view, smoothly!
  sct.scrollIntoView({ block: "end", behavior: "smooth", });
  console.log("Scroll ", id, sct);
};
