<!DOCTYPE html>
<html>
  <head>
    <title>検索結果</title>
    <link rel="stylesheet" type="text/css" href="/assets/css/reset.css" />
    <link rel="stylesheet" type="text/css" href="/assets/css/style.css" />
  </head>
  <body>
    <header class="search-header">
      <form action="/search">
        <input type="text" name="q" class="search-box" value="{{ .Q }}"/>
      </form>
    </header>
    <main class="search-results">
      {{ range .Documents }}
        <div class="search-result">
          <a class="search-result-title" href="{{ .URL }}">{{ .Title }}</a>
          <p class="search-result-description">{{ .Description }}</p>
        </div>
      {{ end }}
    </main>
  </body>
</html>
