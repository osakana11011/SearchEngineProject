<!DOCTYPE html>
<html>
  <head>
    <script src="/assets/js/gtag.js"></script>

    <title>検索結果</title>
    <link rel="stylesheet" type="text/css" href="/assets/css/reset.css" />
    <link rel="stylesheet" type="text/css" href="/assets/css/base.css" />
    <link rel="stylesheet" type="text/css" href="/assets/css/search.css" />

    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="description" content="meta-descriptionです。" />
    <link rel="canonical" href="https://google.com/" />
    <meta name="robots" content="noindex, nofollow">
  </head>
  <body>
    <header class="header">
      <a class="header-title" href="/">
        Wikipedia検索
      </a>
      <form class="header-search-box" action="/search">
        <input type="text" name="q" class="search-box" value="{{ .Q }}"/>
      </form>
    </header>
    <main class="search-results">
      {{ if gt .DocumentsN 0 }}
        {{ range .Documents }}
          <div class="search-result">
            <a class="search-result-title" href="{{ .URL }}">{{ .Title }}</a>
            <p class="search-result-description">{{ .Description }}</p>
          </div>
        {{ end }}
      {{ else }}
        <p>検索結果が見つかりませんでした。</p>
      {{ end }}
    </main>
  </body>
</html>
