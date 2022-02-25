setTimeout(function(){
	document.getElementById("logo").classList.add("visible")
}, 200)

setTimeout(function(){
	document.getElementById("form").classList.add("visible")
}, 750)

function setNick(){
	var nick = document.getElementById("nick").value;
	localStorage.setItem("nick", nick);
}

function showNicks(){
	nicks = localStorage.nicks;
	nicks = nicks.split(",").join("<br/>");
	document.getElementById("nicksDisplay").innerHTML = nicks;
}