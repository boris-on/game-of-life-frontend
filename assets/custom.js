setTimeout(function(){
	document.body.classList.add("visible")
}, 200)

function setNick(){
	var nick = document.getElementById("nick").value;
	localStorage.setItem("nick", nick);
}

function showNicks(){
	nicks = localStorage.nicks;
	nicks = nicks.split(",").join("<br/>");
	document.getElementById("nicksDisplay").innerHTML = nicks;
}