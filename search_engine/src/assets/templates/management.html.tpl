<!DOCTYPE html>
<html>
  <head>
    <title>管理画面</title>
    <link rel="stylesheet" type="text/css" href="/assets/css/reset.css" />
    <link rel="stylesheet" type="text/css" href="/assets/css/base.css" />

    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="description" content="meta-descriptionです。" />
    <link rel="canonical" href="https://google.com/" />
    <meta name="robots" content="noindex, nofollow">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.min.js"></script>
  </head>
  <body>
    <header class="header">
      <a class="header-title" href="/">
        Wikipedia検索
      </a>
    </header>
    <main>
      <div style="width: 600px; height: 400px;">
        <canvas id="documents-graph"></canvas>
      </div>
    </main>

    <script>
      var ctx = document.getElementById("documents-graph");
      var myLineChart = new Chart(ctx, {
        type: 'line',
        data: {
          labels: ['8月1日', '8月2日', '8月3日', '8月4日', '8月5日', '8月6日', '8月7日'],
          datasets: [
            {
              label: '最高気温(度）',
              data: [35, 34, 37, 35, 34, 35, 34, 25],
              borderColor: "rgba(255,0,0,1)",
              backgroundColor: "rgba(0,0,0,0)"
            },
            {
              label: '最低気温(度）',
              data: [25, 27, 27, 25, 26, 27, 25, 21],
              borderColor: "rgba(0,0,255,1)",
              backgroundColor: "rgba(0,0,0,0)"
            }
          ],
        },
        options: {
          title: {
            display: true,
            text: '気温（8月1日~8月7日）'
          },
          scales: {
            yAxes: [{
              ticks: {
                suggestedMax: 40,
                suggestedMin: 0,
                stepSize: 10,
                callback: function(value, index, values){
                  return  value +  '度'
                }
              }
            }]
          },
        }
      });
    </script>
  </body>
</html>
