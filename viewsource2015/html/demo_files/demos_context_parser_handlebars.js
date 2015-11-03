document.addEventListener("DOMContentLoaded", function(event) {

  var preProcessedTemplate = "";
  var templateInput = document.getElementById("templateInput");
  var jsonInput = document.getElementById("jsonInput");

  // templateInput.focus();
  
  function preProcessTemplate() {
    // console.log("preProcessTemplate");
    try {
        var strictMode = document.getElementById("strictMode").checked;
        console.log("[INFO] Strict Mode?"+strictMode);
        var preProcessor = new Handlebars.ContextParserHandlebars({printCharEnable: false, strictMode: strictMode});
        preProcessedTemplate = preProcessor.analyzeContext(templateInput.value);
        document.getElementById("templateOutput").value = preProcessedTemplate;
        return true;
    } catch (err) {
        preProcessedTemplate = err.msg;
        // document.getElementById("htmloutput").innerHTML = err.msg;
    }
  };
  function dataBinding() {
    // console.log("dataBinding");

      try {
          var template = origHandlebars.compile(templateInput.value);
          var data = JSON.parse(jsonInput.value || '{}');
          document.getElementById("origoutput").value = template(data);
          document.getElementById("orightmloutput").innerHTML = template(data);
      } catch (err) {
          document.getElementById("origoutput").value = 'JSON Format ' + err;
      }

      var template = Handlebars.compile(preProcessedTemplate);
      var jsonString = jsonInput.value || '{}';
      try {
          var data = JSON.parse(jsonString);
          document.getElementById("output").value = template(data);
          document.getElementById("htmloutput").innerHTML = template(data);
      } catch (err) {
          document.getElementById("output").value = 'JSON Format ' + err;
      }
  };

  function processTemplate() {
    preProcessTemplate();
    dataBinding();
  }

  templateInput.addEventListener("blur", processTemplate);
  jsonInput.addEventListener("blur", processTemplate);
  jsonInput.addEventListener("keyup", processTemplate);
  processTemplate();

});
