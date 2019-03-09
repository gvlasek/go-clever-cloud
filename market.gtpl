<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <title>Hello!</title>
    <script type="text/javascript">
        getPrices()
        setTimeout(function() {
            getPrices();
        }, 10000);
       function getPrices() {
           fetch('/api')
               .then(resp => resp.json())
               .then(data => {
                   let element = document.getElementById('currencies');

                   string = "<h1>Prices</h1>";
                   for(let s of data){
                        string += s.name + " " + s.price_usd + "<br>";
                   }
                   element.innerHTML = string;
               });
       }
    </script>
</head>
<body>
<h2 id='hello-title'>I am a market</h2>
<div id='currencies'></div>
<button onclick="getPrices()">Refresh</button>
</body>
</html>