function CombinedSplit(s, separator, limit) {
    var arr = s.split(separator, limit);
    var left = s.substring(arr.join(separator).length + separator.length);
    arr.push(left);
    return arr;
}

/**
 * Retrieve information from the JARVIS endpoints periodically
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 * @param {int} everySeconds How often should the API endpoint be polled
 */
function getInfo(elementID, endpointURI, everySeconds) {
    
    // Initial Populate
    _getInfo(elementID, endpointURI);

    if (everySeconds > 0) {
        setInterval(function () {
            _getInfo(elementID, endpointURI);
        }, (everySeconds * 1000));
    } else {
        _getInfo(elementID, endpointURI);
    }
}

/**
 * Retrieve information from JARVIS endpoints 
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 */
function _getInfo(elementID, endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {
               document.getElementById(elementID).innerHTML = xmlhttp.responseText;
           }
        }
    }
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

function HitAPI(endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
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



/**
 * Retrieve information from the JARVIS endpoints periodically
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 * @param {int} everySeconds How often should the API endpoint be polled
 */
function getWorkingOn(elementID, endpointURI, everySeconds) {
    
    // Initial Populate
    _getWorkingOn(elementID, endpointURI);

    if (everySeconds > 0) {
        setInterval(function () {
            _getWorkingOn(elementID, endpointURI);
        }, (everySeconds * 1000));
    } else {
        _getWorkingOn(elementID, endpointURI);
    }
}

/**
 * Retrieve information from JARVIS endpoints 
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 */
function _getWorkingOn(elementID, endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {

                var test = CombinedSplit(xmlhttp.responseText, ",", 1);
                var icon = "";

                // todo output fix check if changed then chang ei t ? 
                if (test[0] == "Bug") {
                    icon = "<img src=\"content/img/jira-bug.svg\" />";
                } else if (test[0] == "New Feature" ) {
                    icon = "<img src=\"content/img/jira-feature.svg\" />";
                } else if (test[0] == "Improvement" ) {
                    icon = "<img src=\"content/img/jira-improvement.svg\" />";
                } else if (test[0] == "Task" ) {
                    icon =  "<img src=\"content/img/jira-task.svg\" />";
                } else if (test[0] == "Sub-Task" ) {
                    icon = "<img src=\"content/img/jira-subtask.svg\" />";
                } else if (test[0] == "Epic" ) {
                    icon = "<img src=\"content/img/jira-epic.svg\" />"
                } 

                if (icon.length > 0 ) {
                    document.getElementById("workingon-image").src = "content/img/jarvis-workon-jira.png"
                    document.getElementById(elementID).innerHTML = icon.concat(test[1]);
                } else {
                    document.getElementById("workingon-image").src = "content/img/jarvis-workon.png"
                    document.getElementById(elementID).innerHTML = xmlhttp.responseText;
                }
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}





/**
 * Retrieve information from the JARVIS endpoints periodically
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 * @param {int} everySeconds How often should the API endpoint be polled
 */
function getJIRA(elementID, endpointURI, renderElement, everySeconds) {
    
    // Initial Populate
    _getJIRA(elementID, endpointURI, renderElement);

    if (everySeconds > 0) {
        setInterval(function () {
            _getJIRA(elementID, endpointURI, renderElement);
        }, (everySeconds * 1000));
    } else {
        _getJIRA(elementID, endpointURI, renderElement);
    }
}

/**
 * Retrieve information from JARVIS endpoints 
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 */
function _getJIRA(elementID, endpointURI, renderElement) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {

                var returnValue = "";
                var elements = xmlhttp.responseText.split("\n");
                for (var i = 0, len = elements.length; i < len; i++) {
                    
                    var test = CombinedSplit(elements[i], ",", 2);
                    var icon = "";

                    // todo output fix check if changed then chang ei t ? 
                    if (test[1] == "Bug") {
                        icon = "<img src=\"content/img/jira-bug.svg\" />";
                    } else if (test[1] == "New Feature" ) {
                        icon = "<img src=\"content/img/jira-feature.svg\" />";
                    } else if (test[1] == "Improvement" ) {
                        icon = "<img src=\"content/img/jira-improvement.svg\" />";
                    } else if (test[1] == "Task" ) {
                        icon =  "<img src=\"content/img/jira-task.svg\" />";
                    } else if (test[1] == "Sub-Task" ) {
                        icon = "<img src=\"content/img/jira-subtask.svg\" />";
                    } else if (test[1] == "Epic" ) {
                        icon = "<img src=\"content/img/jira-epic.svg\" />"
                    } 

                    if (icon.length > 0 ) {
                        returnValue = returnValue + renderElement(icon,test[0], test[2]);
                    } else {
                        returnValue = returnValue + "<p>" + elements[i] + "</p>";
                       
                    }
                }



                document.getElementById(elementID).innerHTML = returnValue;
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}


