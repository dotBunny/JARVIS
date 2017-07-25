/**
 * Retrieve information from the JARVIS endpoints periodically
 * @param {string} endpointURI The full URI of the API endpoint to poll
 * @param {function} callbackFunction The function to callback with the dataset
 * @param {int} everySeconds How often should the API endpoint be polled
 */
function getJSON(endpointURI,callbackFunction, everySeconds) {
    
    var responseHandler = callbackFunction;
    _getJSON(endpointURI, responseHandler);
    if (everySeconds > 0) {
        setInterval(function () {
            _getJSON(endpointURI, responseHandler);
        }, (everySeconds * 1000));
    } else {
        _getJSON(endpointURI, responseHandler);
    }
}

/**
 * Retrieve information from JARVIS endpoints 
 * @param {function} callbackFunction The function to callback with the dataset
 * @param {string} endpointURI The full URI of the API endpoint to poll
 */
function _getJSON(endpointURI, callbackFunction) {
    var xmlhttp = new XMLHttpRequest();
    var responseHandler = callbackFunction;

    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState == XMLHttpRequest.DONE && xmlhttp.status == 200 && responseHandler) {
            responseHandler(JSON.parse(xmlhttp.responseText));
        }
    };

    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
    
}




function refreshImage(elementID, everySeconds) {

    // Save base source    
    document.getElementById(elementID).setAttribute("data-source", document.getElementById(elementID).src);
    setInterval(function () { 
        d = new Date();
        var newSource = document.getElementById(elementID).getAttribute("data-source") + "?" + d.getTime()        
        document.getElementById(elementID).src = newSource;
    }, (everySeconds * 1000));
}

function hitAPI(endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}

















function CombinedSplit(s, separator, limit) {
    var arr = s.split(separator, limit);
    var left = s.substring(arr.join(separator).length + separator.length);
    arr.push(left);
    return arr;
}










function getList(elementID, endpointURI, renderElement, everySeconds) {
    
    // Initial Populate
    _getList(elementID, endpointURI, renderElement);

    if (everySeconds > 0) {
        setInterval(function () {
            _getList(elementID, endpointURI, renderElement);
        }, (everySeconds * 1000));
    } else {
        _getList(elementID, endpointURI,renderElement);
    }
}
function _getList(elementID, endpointURI, renderElement) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {
                var items = xmlhttp.responseText.split(",");
                var output = ""
                for (var i = 0; i < items.length; i++) {
                    output = output.concat(renderElement(items[i]))
                }
               document.getElementById(elementID).innerHTML = output
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}



var LastMediaCount = 0;
function StartMonitoringMedia(endpointURI) { 
    _monitorMedia(endpointURI, false);
    setInterval(function () {_monitorMedia(endpointURI, true);}, (1 * 1000)); }

function _monitorMedia(endpointURI, notify)
{
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {
                var items = xmlhttp.responseText.split(",");
                var count = parseInt(items[0]);
                
                if ( count > LastMediaCount && count > 0) {
                    if ( notify ) {
                        var audio = new Audio(items[1]);
                        audio.play();
                    }
                    LastMediaCount = count;
                }
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}