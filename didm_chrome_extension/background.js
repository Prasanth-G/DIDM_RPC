var PeerIPs = [];
var NoOfParts = 8;
var override = false;

function startDownload(url) {

    xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:9090/", true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 202){
                var notification = new Notification('Download Complete', {
                    icon: 'images/logo.png',
                    body: "Hey there! Your download was Ready.\n Wanna Open ?",
                });
                notification.onclick = function () {
                    chrome.downloads.showDefaultFolder();
                };
            } else if (xhr.status == 0){
              alert("Unable to establish connection to server");
            } else {
              alert(`Status: ${xhr.status}; ${xhr.responseText}`);
            }
        }
    }
    var data = JSON.stringify({
        "Url" : url,
        "PeerIPAddr" : JSON.stringify(PeerIPs),
        "NoOfParts" : NoOfParts
    });
    alert("Download Link: " + url + '\n Parts: ' + NoOfParts);
    xhr.send(data);
}

function overrideDownload(downloaditem){
    if (override){
        url = (' ' + downloaditem.finalUrl).slice(1);
        chrome.downloads.cancel(downloaditem.id);
        //chrome.downloads.erase(downloaditem.id, function(){
        //    alert("CAnt");
        //});

        startDownload(url);
    }
    //alert("DIDM not set as default");
}

function downloadSelectedLink(ctxitem){
    startDownload(ctxitem.linkUrl);
}

window.onload = function(){
    chrome.downloads.onCreated.addListener(overrideDownload);
    chrome.contextMenus.create({
        title: "Download with DIDM", 
        contexts:["link"], 
        onclick: downloadSelectedLink
    });
}