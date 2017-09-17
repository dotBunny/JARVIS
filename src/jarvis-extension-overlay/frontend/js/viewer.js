var bIsPanelOpen = false;
var bIsShowingStatsPanel = false;
var bIsShowingJIRAPanel = false;
var bIsShowingSpotifyPanel = false

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
  $.getJSON( "https://api.dotbunny.com/v1/JARVIS/Poll", function( data ) {
  
  console.log(data);
  // var items = [];
  // $.each( data, function( key, val ) {
  //   items.push( "<li id='" + key + "'>" + val + "</li>" );
  // });
 
  // $( "<ul/>", {
  //   "class": "my-new-list",
  //   html: items.join( "" )
  // }).appendTo( "body" );
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



var ctx = document.getElementById("myChart").getContext('2d');
var myChart = new Chart(ctx, {
    type: 'horizontalBar',
    data: {
        labels: ["Crashes", "Saves", "Swears", "Coffee", "Builds", "Orange"],
        datasets: [{
            label: '# of Votes',
            data: [12, 19, 3, 5, 2, 3],
            backgroundColor: [
                'rgba(255, 99, 132, 0.2)',
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(153, 102, 255, 0.2)',
                'rgba(255, 159, 64, 0.2)'
            ],
            borderColor: [
                'rgba(255,99,132,1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(255, 159, 64, 1)'
            ],
            borderWidth: 1
        }]
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