
var sock = null;
var wsuri = "ws://s1.natapp.cc:54888/server_client/";


var  Date1 = new Array() 
var Price = new Array()
var Change = new Array()
var Orange=new Array()
 
window.onload = function() {	

    //document.getElementById("CrossSectionalArea").onkeyup= ExamineCrossSectionalArea();
    sock = new WebSocket(wsuri);
    sock.onopen = function() {
        console.log("connected to " + wsuri);
    }
    sock.onclose = function(e) {
        console.log("connection closed (" + e.code + ")");
    }
    sock.onmessage = function(e) {
        //console.log("message received: " + e.data);
        //show(e.data)
        //update1(e.data)
    }
    
};

function send() {
    var msg ="========="
    sock.send(msg);
};

function Draw(myChart,date){

    var i =0
    
    var da = date.split("|");
    for (let index = 0; index < da.length-1; index=index+4) {
        Orange[i]=da[index]
        Price[i]= da[index +1]
        Change[i]=da[index+2]
        Date1[i] = da[index +3]
       i++
    }
   // myChart.data.labels=Date1;
    //myChart.data.datasets[0]=Price;
    myChart.update()
}