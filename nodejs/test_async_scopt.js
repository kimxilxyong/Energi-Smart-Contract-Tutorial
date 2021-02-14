#! /usr/bin/env node

const arrayItems = [];
const execute = async (count) => {
    /* let item; */
    for (let i = 0; i < count; ++i) {
        let item = "item" + i;
        arrayItems.push(i);

        setInterval(() => {
            if (item === "item1") {
                i = i * 2;
            } else {
                i++;
            }
            console.log(count--, item, i, getTime());
        }, 2000);
    }
    for (let i = 0; i < count; ++i) {
        let item = "Ditem" + i;
        arrayItems.push(i);

        setInterval(() => {
            console.log(count--, item, i--, getTime());
        }, 2000);
    }
    // Monitor func
    setInterval(() => {
            console.log(arrayItems);
        }, 9000);

};

function getTime() {

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
}

execute(2);
