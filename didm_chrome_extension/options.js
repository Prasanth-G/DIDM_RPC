function addPeerInputBox(){
  var container = document.getElementById("p1");
  var a = chrome.downloads.showDefaultFolder();
  container.innerHTML = "SUPPI";
  container.nodeValue = "VA";
  alert()
  // container.appendChild(document.createTextNode("Member " + (i+1)));
  // // Create an <input> element, set its type and name attributes
  // var input = document.createElement("input");
  // input.type = "text";
  // input.name = "member" + i;
  // container.appendChild(input);
  // // Append a line break 
  // container.appendChild(document.createElement("br"));
}

window.onload = function() {
  document.getElementById('AddPeers').onclick = addPeerInputBox;
}
