import { writable } from 'svelte/store';

let visitedSvelte = {
    info: false,
    wallet: false,
    web3: false,
    smartcontract: false,
  };

export const global = writable(visitedSvelte);


export const getTime = () => {
  let str = "";

  let currentTime = new Date()
  let hours = currentTime.getHours()
  let minutes = currentTime.getMinutes()
  let seconds = currentTime.getSeconds()

  if (minutes < 10) {
      minutes = "0" + minutes
  }
  if (seconds < 10) {
      seconds = "0" + seconds
  }
  str += hours + ":" + minutes + ":" + seconds + " ";
  if(hours > 11){
      str += "PM"
  } else {
      str += "AM"
  }
  return str;
}
