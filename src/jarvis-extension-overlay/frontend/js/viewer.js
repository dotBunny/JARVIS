var bIsPanelOpen = false;
var bIsShowingStatsPanel = false;
var bIsShowingJIRAPanel = false;
var bIsShowingSpotifyPanel = false

var spotifyLink;
var spotifyThumbnail;
var spotifyTrack;
var spotifyArtist;
var jiraList;
var jiraLatestID;
var statsChart = document.getElementById("stats").getContext('2d');
var createdChart = false;
var statsChartObject;

var pollingHandle;





// Icon 
$("li#button-spotify").on("click", function() {
  if ( !bIsShowingSpotifyPanel ) 
  {
    if ( !bIsPanelOpen) { openPanel(); }
    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingSpotifyPanel = true;
    $("div#panel-spotify").animate({right: "0px"});
    bIsShowingJIRAPanel = false;
    $("div#panel-jira").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingStatsPanel = false;
    $("div#panel-stats").animate({right: String(-((window.innerWidth/100)*30)) + "px"});   
  } 
  else 
  {
    bIsShowingSpotifyPanel = false;
    closePanel();
  }
});


$("li#button-stats").on("click", function() {
  if ( !bIsShowingStatsPanel ) 
  {
    if ( !bIsPanelOpen) { openPanel(); }
    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingStatsPanel = true;
    $("div#panel-stats").animate({right: "0px"});
    bIsShowingSpotifyPanel = false;
    $("div#panel-spotify").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingJIRAPanel = false;
    $("div#panel-jira").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
  } 
  else 
  {
    bIsShowingStatsPanel = false;
    closePanel();
  }
});

$("li#button-jira").on("click", function() {
  if ( !bIsShowingJIRAPanel ) 
  {
    if ( !bIsPanelOpen) { openPanel(); }
    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingJIRAPanel = true;
    $("div#panel-jira").animate({right: "0px"});
    bIsShowingSpotifyPanel = false;
    $("div#panel-spotify").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingStatsPanel = false;
    $("div#panel-stats").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    
  } 
  else 
  {
    bIsShowingJIRAPanel = false;
    closePanel();
  }
});



$("div#cancel").on("click", function() {
  if ( bIsPanelOpen ) {
    closePanel();
  }
});




function openPanel()
{
  $("div#jarvis").animate({right: "0px"});
  bIsPanelOpen = true;
  poll();
  clearInterval(pollingHandle);
  pollingHandle = setInterval(function () {poll(); }, 30000);
}
function closePanel(animate = true)
{
  var amount = String(-((window.innerWidth/100)*30)) + "px";

  if ( animate ) {
    $("div#panel-spotify").animate({right: amount});
    $("div#panel-jira").animate({right: amount});
    $("div#panel-stats").animate({right: amount});
    $("div#jarvis").animate({right: amount});
  } else {
    $("div#panel-spotify").css("right", amount);
    $("div#panel-jira").css("right", amount);
    $("div#panel-stats").css("right", amount);
    $("div#jarvis").css("right", amount); 
  }
  $(".active").removeClass("active");

  bIsPanelOpen = false;
  bIsShowingJIRAPanel = false;
  bIsShowingSpotifyPanel = false;
  bIsShowingStatsPanel = false;

  clearInterval(pollingHandle);
  pollingHandle = setInterval(function () {poll(); }, 300000);
}



function poll() {
  $.getJSON("https://api.dotbunny.com/v1/JARVIS/Poll", function (data) {
    

    // Is a different track
    if ($(spotifyLink).attr('href') != data['spotify']['CurrentlyPlayingURL']) {
      $(spotifyLink).attr('href', data['spotify']['CurrentlyPlayingURL']);

      $(spotifyArtist).html(data['spotify']['ArtistLine']);
      $(spotifyTrack).html(data['spotify']['TrackName']);

      $(spotifyThumbnail).attr('src', data['spotify']['TrackThumbnailURL']);
    }


    // Process JIRA
    if ( data['tasks']['Mode'] == "JIRA") 
    {
      var items = Array();

      console.log(data);
      if ( data['tasks']['List'][0]['ID'] != jiraLatestID ) {

        jiraLatestID = data['tasks']['List'][0]['ID'];
        // Clear out the list
        $(jiraList).empty();

        var itemCount = data['tasks']['List'].length;
        for (var i = 0; i < itemCount; i++) 
        {
          var item = '<li>';
          switch(data['tasks']['List'][i]['Type'])
          {
            case "bug":
              item += '<img src="img/task-bug.png" />';
              break;
            case "epic":
              item += '<img src="img/task-epic.png" />';
              break;
            case "improvement":
              item += '<img src="img/task-improvement.png" />';
              break;
            case "task":
              item += '<img src="img/task-task.png" />';
              break;
            case "new":
            default:
              item += '<img src="img/task-new.png" />';
              break;
          }

          item += '<p>' + data['tasks']['List'][i]['Title'];
          item += '</p></li>';

          $(jiraList).append(item);
        }
      }

      
    }

    // Process Stats Array
    var newDataObject = {};   
    
    newDataObject["data"] = Array();
    newDataObject["backgroundColor"] = Array();
    newDataObject["borderColor"] = Array();
    newDataObject["borderWidth"] = 1;
    
    Object.keys(data['stats']).forEach(function (index) {
      newDataObject["data"].push(data['stats'][index]['CurrentValue']);
      newDataObject["backgroundColor"].push(hexToRgbA(data['stats'][index]['Color'], "0.2"));
      newDataObject["borderColor"].push(hexToRgbA(data['stats'][index]['Color'], "1"));
    });
    


    if (createdChart) {  
      statsChartObject.data.datasets[0] = newDataObject;
      statsChartObject.update(0);  
    } else {

      statsChartObject = new Chart(statsChart, {
        type: 'horizontalBar',
        data: {
          labels: Object.keys(data['stats']),
          datasets: [newDataObject]
        },
        options: {
            legend: {
              display: false
            },
            tooltips: {
              enabled: false
            },
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero:true
                    }
                }]
            }
        }
      });
      createdChart = true;  
    }
  });
}


// Helpers
function updateSquares()
{

  var containerWidth = $(".square-container").width();

  var squareWidth = (containerWidth / 3) - 5;
  $(".square").each(function(index) {
    $(this).css("height", squareWidth + 5);
    $(this).css("width", squareWidth);

    $(this).children("i").css("line-height", (squareWidth + 5)+'px');
    $(this).children("p.count").css("line-height", (squareWidth + 5)+'px');
  });
}

// Events
$( window ).resize(function() {
  closePanel(false);
  updateSquares();
});




// Startup
poll();
closePanel();
updateSquares();




$( document ).ready(function() {
  
  // Spotify References
  spotifyLink = $('div#panel-spotify a#spotify-link');
  spotifyThumbnail = $('div#panel-spotify img#spotify-thumbnail');
  spotifyTrack = $('div#panel-spotify p#spotify-track');
  spotifyArtist = $('div#panel-spotify p#spotify-artist');

  // JIRA References
  jiraList = $('div#panel-jira ul');
});

function hexToRgbA(hex, opacity){
  var c;
  if(/^#([A-Fa-f0-9]{3}){1,2}$/.test(hex)){
      c= hex.substring(1).split('');
      if(c.length== 3){
          c= [c[0], c[0], c[1], c[1], c[2], c[2]];
      }
      c= '0x'+c.join('');
      return 'rgba('+[(c>>16)&255, (c>>8)&255, c&255].join(',')+',' + opacity +')';
  }
  throw new Error('Bad Hex');
}