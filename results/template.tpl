<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Dorkscout results</title>
    <meta name="description" content="Dorkscout results">
    <meta name="author" content="Dorkscout">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!--[if lt IE 9]>
<script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
<![endif]-->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">
    <style>
      body {
        margin: 1.5em auto;
        font-family: "Helvetica Neue", sans-serif;
        color: #575553;
        max-width: 1000px;
        padding: 0 1em;
      }

      h1 {
        font-size: 1.75em;
        color: #003e54;
        font-weight: bold;
        line-height: 1.4;
        margin: 0 0 0.8em 0;
      }

      .item {
        padding: 1em;
        border-bottom: 1px solid #f0f0f0;
      }

      .dorkscout-item {
        background-color: #f0f0f0;
      }

      p {
        font-size: 1em;
        line-height: 1.6;
        margin: 0 0 0.3em 0;
      }

      p:last-of-type {
        margin: 0;
      }

      span.divider-dot {
        margin: 0 0.3em;
        color: #999794;
      }

      .handle {
        font-weight: bold;
        color: #000000;
      }

      a {
        color: #2d9eb2;
        text-decoration: none;
      }

      a:active,
      a:hover {
        color: #207180;
      }

      img {
        border-radius: 4px;
      }
    </style>
    <meta name="user-style-sheet" content="pdf.css">
  </head>
  <body>
    <h1 style="color:#000000">Dorkscout results {{.Target}}</h1>
    {{.Time}}
    {{.Next}}
  </body>
</html>