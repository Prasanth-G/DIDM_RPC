function initiateDownload(){
  url = document.getElementById("dinput").value;
  if (!isValidURL(url)){
    alert("Enter a Valid URL");
    return;
  }
  background.startDownload(url);
}


function isValidURL(url){
  url = url.trim();
  var res = url.match(/(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g);
  if(res == null)
    return false;
  else
    return true;
}

function isValidIP(ip){
  s = ip.split(".");
  valid = true;
  number = 0;
  for (i = 0; i < s.length; i++){
    e = parseInt(s[i], 10);
    if (number++ > 3){
      valid = false;
      break;
    }
    if (Number.isNaN(e)){
      valid = false;
      break;
    }
    if (e < 0 && e > 255){
      valid = false;
      break;
    }
  }
  return valid;
}

function delme(event){

}

function addPeerInputBox(){
  ip = document.getElementById("ipInput").value;
  if (isValidIP(ip)) {
    ip = ip + DefaultRPCServerPort;
    if (background.PeerIPs.indexOf(ip) == -1){
      //add to list of devices
      AddPeerIPtoStorage(ip);
      AddPeerIPtoUI(ip);
    } else {
      alert("IP Already Exists")
    }
  
  }else {
    alert("Enter a Valid IP. \n Eg: 192.168.0.1")
  }
}

function AddPeerIPtoStorage(ip){
  background.PeerIPs.push(ip);
}

function updateNoofParts(event){
  background.NoOfParts = document.getElementById('NoOfParts').value;
}

function removeAllPeers(){
  background.PeerIPs = [];
  UpdatePeersUI();
}

function UpdatePeersUI(){
  peerList = document.getElementById("peerIPList");
  while (peerList.hasChildNodes()) {   
    peerList.removeChild(peerList.firstChild);
  }
  background.PeerIPs.forEach(function (key){
    AddPeerIPtoUI(key);
  });
}

function toggleDefaultFunc(){
  background.override = document.getElementById('toggleDefault').checked;
  UpdateSetDefaultUI();
}

function UpdateSetDefaultUI(){
  if (background.override){
    document.getElementById("downloadarea").style.display = "none";
  } else {
    document.getElementById("downloadarea").style.display = "block";
  }
}

var background = chrome.extension.getBackgroundPage();
var DefaultRPCServerPort = ":50051";

window.onload = function() {
  document.getElementById('startDownloadbtn').onclick = initiateDownload;
  document.getElementById('toggleSettings').onclick = showSettings;
  document.getElementById('AddPeers').onclick = addPeerInputBox;
  document.getElementById('removePeers').onclick = removeAllPeers;
  np = document.getElementById('NoOfParts');
  np.onchange = updateNoofParts;
  np.value = background.NoOfParts;
  tg = document.getElementById('toggleDefault');
  tg.onchange = toggleDefaultFunc;
  tg.checked = background.override;
  UpdateSetDefaultUI();
  
  document.getElementById("settings").style.display = "none";
  
  UpdatePeersUI();
};

//For UI Changes
function AddPeerIPtoUI(ip) {
  peerList = document.getElementById("peerIPList");
  var number = document.getElementById('peerIPList').getElementsByTagName('input').length + 1;
  peerList.appendChild(document.createTextNode("Peer " + number + " : "));
  var input = document.createElement("input");
  input.className = "ip";
  input.id = number;
  input.readOnly = true;
  input.value = ip;
  delbtn = document.createElement("button");
  delbtn.className = "xbtn";
  delbtn.id = number;
  delbtn.onclick = delme;
  delbtn.appendChild(document.createTextNode("X"));
  input.appendChild(delbtn);
  peerList.appendChild(input);
  peerList.appendChild(document.createElement("br"));
}

function showSettings() {
  var x = document.getElementById("settings");
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}
