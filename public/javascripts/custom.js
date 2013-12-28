$(document).ready(function () {
    var colors = ["#fb2900", "#ff7800", "#fff43b", "#8dfa30", "#01c86a", "#00d7b2", "#0092e3", "#002f7e", "#390e73"];
    setInterval(function () {
        $("body").css("background-color",colors[Math.floor(Math.random() * colors.length)]);
        
    }, 5000);
});