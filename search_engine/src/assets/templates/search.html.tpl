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
    <meta name="robots" content="noindex, nofollow">
  </head>
  <body>
    <header class="header">
      <a class="header-title is-pc" href="/">
        Wiki検索
      </a>
      <form action="/search" method="get">
        <input class="search-box" type="text" name="q" value="{{ .Q }}"/>
      </form>
    </header>

    <main class="main">
      {{ if gt .DocumentsN 0 }}
      <div class="search-results">
        {{ range .Documents }}
          <div class="search-result">
            <a class="search-result-title" href="{{ .URL }}">{{ .Title }}</a>
            <p class="search-result-description">{{ .Description }}</p>
          </div>
        {{ end }}

        <div class="paging">
          <ul class="paging-numbers">
            {{ range .Pages }}
              {{ if .IsCurrent }}
                <li class="paging-number current"><a href="{{ .URL }}">{{ .Number }}</a></li>
              {{ else }}
                <li class="paging-number"><a href="{{ .URL }}">{{ .Number }}</a></li>
              {{ end }}
            {{ end }}
          </ul>
        </div>
      </div>
      {{ else }}
        <p>検索結果が見つかりませんでした。</p>
      {{ end }}
    </main>

    <header class="footer">
    </header>
  </body>
</html>
