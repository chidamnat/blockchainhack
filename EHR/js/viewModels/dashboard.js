/**
 * Copyright (c) 2014, 2017, Oracle and/or its affiliates.
 * The Universal Permissive License (UPL), Version 1.0
 */
/*
 * Your dashboard ViewModel code goes here
 */
define(['ojs/ojcore', 'knockout', 'jquery','ojs/ojarraytabledatasource','ojs/ojtable','ojs/ojchart','ojs/ojpopup'],
 function(oj, ko, $) {
  
    function DashboardViewModel() {
      var self = this;
      
      // Below are a subset of the ViewModel methods invoked by the ojModule binding
      // Please reference the ojModule jsDoc for additionaly available methods.

      /**
       * Optional ViewModel method invoked when this ViewModel is about to be
       * used for the View transition.  The application can put data fetch logic
       * here that can return a Promise which will delay the handleAttached function
       * call below until the Promise is resolved.
       * @param {Object} info - An object with the following key-value pairs:
       * @param {Node} info.element - DOM element or where the binding is attached. This may be a 'virtual' element (comment node).
       * @param {Function} info.valueAccessor - The binding's value accessor.
       * @return {Promise|undefined} - If the callback returns a Promise, the next phase (attaching DOM) will be delayed until
       * the promise is resolved
       */
      self.handleActivated = function(info) {
        // Implement if needed
      };

      /**
       * Optional ViewModel method invoked after the View is inserted into the
       * document DOM.  The application can put logic that requires the DOM being
       * attached here.
       * @param {Object} info - An object with the following key-value pairs:
       * @param {Node} info.element - DOM element or where the binding is attached. This may be a 'virtual' element (comment node).
       * @param {Function} info.valueAccessor - The binding's value accessor.
       * @param {boolean} info.fromCache - A boolean indicating whether the module was retrieved from cache.
       */
      self.handleAttached = function(info) {
        // Implement if needed
      };


      /**
       * Optional ViewModel method invoked after the bindings are applied on this View. 
       * If the current View is retrieved from cache, the bindings will not be re-applied
       * and this callback will not be invoked.
       * @param {Object} info - An object with the following key-value pairs:
       * @param {Node} info.element - DOM element or where the binding is attached. This may be a 'virtual' element (comment node).
       * @param {Function} info.valueAccessor - The binding's value accessor.
       */
      self.handleBindingsApplied = function(info) {
        // Implement if needed
      };

      /*
       * Optional ViewModel method invoked after the View is removed from the
       * document DOM.
       * @param {Object} info - An object with the following key-value pairs:
       * @param {Node} info.element - DOM element or where the binding is attached. This may be a 'virtual' element (comment node).
       * @param {Function} info.valueAccessor - The binding's value accessor.
       * @param {Array} info.cachedNodes - An Array containing cached nodes for the View if the cache is enabled.
       */
      self.handleDetached = function(info) {
        // Implement if needed
      };
    }

    /*
     * Returns a constructor for the ViewModel so that the ViewModel is constrcuted
     * each time the view is displayed.  Return an instance of the ViewModel if
     * only one instance of the ViewModel is needed.
     */
    
    var claims =[
                        {"claimId":"100001","patientId":"0001","npId":"1001","cptId":"001","icdId":"01","ndc":"Med1,Med2","cost":"100","colorCode":"black"},
                        {"claimId":"100002","patientId":"0002","npId":"1001","cptId":"002","icdId":"02","ndc":"Med3,Med2","cost":"200","colorCode":"black"},
                        {"claimId":"100003","patientId":"0003","npId":"1002","cptId":"003","icdId":"03","ndc":"Med1,Med4","cost":"400","colorCode":"black"},
                        {"claimId":"100004","patientId":"0004","npId":"1001","cptId":"001","icdId":"01","ndc":"Med1,Med5","cost":"100","colorCode":"black"},
                        {"claimId":"100005","patientId":"0005","npId":"1003","cptId":"002","icdId":"02","ndc":"Med6,Med2","cost":"500","colorCode":"red"},
                        {"claimId":"100006","patientId":"0006","npId":"1001","cptId":"001","icdId":"01","ndc":"Med9,Med2","cost":"300","colorCode":"red"},
                        {"claimId":"100007","patientId":"0007","npId":"1001","cptId":"003","icdId":"03","ndc":"Med7,Med2","cost":"800","colorCode":"red"}
                ];
    var claims2 = [{"cptId":5,"claimId":1,"patientId":8,"npId":10001,"cost":476,"ndc":19,"icdId":8,"date":"4/20/2017"},
{"cptId":2,"claimId":2,"patientId":3,"npId":10005,"cost":498,"ndc":8,"icdId":11,"date":"4/14/2017"},
{"cptId":8,"claimId":3,"patientId":24,"npId":10006,"cost":445,"ndc":3,"icdId":18,"date":"4/30/2017"},
{"cptId":4,"claimId":4,"patientId":26,"npId":10011,"cost":413,"ndc":16,"icdId":5,"date":"4/17/2017"},
{"cptId":3,"claimId":5,"patientId":27,"npId":10016,"cost":485,"ndc":2,"icdId":14,"date":"4/8/2017"},
{"cptId":5,"claimId":6,"patientId":7,"npId":10023,"cost":537,"ndc":4,"icdId":16,"date":"4/7/2017"},
{"cptId":5,"claimId":7,"patientId":7,"npId":10029,"cost":476,"ndc":5,"icdId":18,"date":"4/29/2017"},
{"cptId":1,"claimId":8,"patientId":18,"npId":10033,"cost":586,"ndc":3,"icdId":18,"date":"4/1/2017"},
{"cptId":5,"claimId":9,"patientId":24,"npId":10039,"cost":523,"ndc":9,"icdId":17,"date":"4/4/2017"},
{"cptId":6,"claimId":10,"patientId":22,"npId":10040,"cost":439,"ndc":18,"icdId":10,"date":"4/7/2017"},
{"cptId":2,"claimId":11,"patientId":5,"npId":10046,"cost":449,"ndc":20,"icdId":18,"date":"4/25/2017"},
{"cptId":3,"claimId":12,"patientId":16,"npId":10055,"cost":441,"ndc":18,"icdId":16,"date":"4/30/2017"},
{"cptId":1,"claimId":13,"patientId":3,"npId":10056,"cost":419,"ndc":19,"icdId":17,"date":"4/8/2017"},
{"cptId":1,"claimId":14,"patientId":7,"npId":10078,"cost":438,"ndc":9,"icdId":4,"date":"4/21/2017"},
{"cptId":10,"claimId":15,"patientId":1,"npId":10083,"cost":438,"ndc":7,"icdId":7,"date":"4/17/2017"},
{"cptId":6,"claimId":16,"patientId":28,"npId":10085,"cost":451,"ndc":6,"icdId":11,"date":"4/21/2017"},
{"cptId":3,"claimId":17,"patientId":2,"npId":10090,"cost":397,"ndc":11,"icdId":3,"date":"4/15/2017"},
{"cptId":9,"claimId":18,"patientId":11,"npId":10092,"cost":518,"ndc":2,"icdId":7,"date":"4/14/2017"},
{"cptId":10,"claimId":19,"patientId":29,"npId":10100,"cost":390,"ndc":14,"icdId":4,"date":"4/17/2017"},
{"cptId":3,"claimId":20,"patientId":25,"npId":10103,"cost":496,"ndc":1,"icdId":4,"date":"4/2/2017"},
{"cptId":6,"claimId":21,"patientId":10,"npId":10104,"cost":447,"ndc":12,"icdId":1,"date":"4/7/2017"},
{"cptId":10,"claimId":22,"patientId":29,"npId":10113,"cost":422,"ndc":14,"icdId":7,"date":"4/20/2017"},
{"cptId":2,"claimId":23,"patientId":1,"npId":10139,"cost":394,"ndc":7,"icdId":9,"date":"4/12/2017"},
{"cptId":10,"claimId":24,"patientId":18,"npId":20017,"cost":641,"ndc":12,"icdId":6,"date":"4/6/2017"},
{"cptId":1,"claimId":25,"patientId":1,"npId":30002,"cost":695,"ndc":16,"icdId":6,"date":"4/3/2017"},
{"cptId":5,"claimId":26,"patientId":23,"npId":30006,"cost":576,"ndc":16,"icdId":9,"date":"4/12/2017"},
{"cptId":3,"claimId":27,"patientId":18,"npId":30007,"cost":801,"ndc":18,"icdId":18,"date":"4/27/2017"},
{"cptId":2,"claimId":28,"patientId":15,"npId":30010,"cost":538,"ndc":6,"icdId":12,"date":"4/25/2017"},
{"cptId":5,"claimId":29,"patientId":8,"npId":30011,"cost":490,"ndc":12,"icdId":1,"date":"4/26/2017"},
{"cptId":4,"claimId":30,"patientId":29,"npId":30012,"cost":613,"ndc":5,"icdId":12,"date":"4/10/2017"},
{"cptId":1,"claimId":31,"patientId":8,"npId":30013,"cost":597,"ndc":6,"icdId":18,"date":"4/7/2017"},
{"cptId":5,"claimId":32,"patientId":29,"npId":30023,"cost":906,"ndc":15,"icdId":13,"date":"4/4/2017"},
{"cptId":2,"claimId":33,"patientId":13,"npId":30024,"cost":742,"ndc":20,"icdId":2,"date":"4/7/2017"},
{"cptId":7,"claimId":34,"patientId":10,"npId":30030,"cost":654,"ndc":4,"icdId":6,"date":"4/5/2017"},
{"cptId":1,"claimId":35,"patientId":12,"npId":30038,"cost":547,"ndc":3,"icdId":2,"date":"4/25/2017"},
{"cptId":10,"claimId":36,"patientId":28,"npId":30061,"cost":495,"ndc":1,"icdId":8,"date":"4/13/2017"},
{"cptId":9,"claimId":37,"patientId":14,"npId":30062,"cost":732,"ndc":11,"icdId":2,"date":"4/10/2017"},
{"cptId":7,"claimId":38,"patientId":2,"npId":30065,"cost":485,"ndc":20,"icdId":10,"date":"4/26/2017"},
{"cptId":6,"claimId":39,"patientId":9,"npId":30069,"cost":567,"ndc":9,"icdId":19,"date":"4/5/2017"},
{"cptId":10,"claimId":40,"patientId":12,"npId":30087,"cost":517,"ndc":7,"icdId":16,"date":"4/3/2017"},
{"cptId":5,"claimId":41,"patientId":29,"npId":30089,"cost":585,"ndc":5,"icdId":2,"date":"4/24/2017"},
{"cptId":4,"claimId":42,"patientId":18,"npId":30092,"cost":520,"ndc":17,"icdId":8,"date":"4/24/2017"},
{"cptId":2,"claimId":43,"patientId":28,"npId":30093,"cost":531,"ndc":15,"icdId":12,"date":"4/9/2017"},
{"cptId":2,"claimId":44,"patientId":7,"npId":30094,"cost":626,"ndc":19,"icdId":15,"date":"4/5/2017"},
{"cptId":1,"claimId":45,"patientId":20,"npId":30100,"cost":359,"ndc":11,"icdId":18,"date":"4/6/2017"},
{"cptId":9,"claimId":46,"patientId":18,"npId":30101,"cost":565,"ndc":11,"icdId":5,"date":"4/23/2017"},
{"cptId":1,"claimId":47,"patientId":13,"npId":30103,"cost":567,"ndc":1,"icdId":7,"date":"4/17/2017"},
{"cptId":9,"claimId":48,"patientId":29,"npId":30105,"cost":479,"ndc":8,"icdId":1,"date":"4/21/2017"},
{"cptId":1,"claimId":49,"patientId":25,"npId":40004,"cost":437,"ndc":2,"icdId":10,"date":"4/14/2017"},
{"cptId":3,"claimId":50,"patientId":19,"npId":40007,"cost":484,"ndc":7,"icdId":10,"date":"4/9/2017"},
{"cptId":6,"claimId":51,"patientId":12,"npId":40010,"cost":458,"ndc":14,"icdId":9,"date":"4/9/2017"},
{"cptId":7,"claimId":52,"patientId":7,"npId":40020,"cost":525,"ndc":17,"icdId":7,"date":"4/25/2017"},
{"cptId":6,"claimId":53,"patientId":3,"npId":40022,"cost":423,"ndc":4,"icdId":18,"date":"4/2/2017"},
{"cptId":6,"claimId":54,"patientId":17,"npId":40026,"cost":476,"ndc":15,"icdId":19,"date":"4/24/2017"},
{"cptId":2,"claimId":55,"patientId":24,"npId":40027,"cost":442,"ndc":6,"icdId":16,"date":"4/28/2017"},
{"cptId":6,"claimId":56,"patientId":15,"npId":40029,"cost":457,"ndc":8,"icdId":4,"date":"4/15/2017"},
{"cptId":1,"claimId":57,"patientId":12,"npId":40036,"cost":462,"ndc":13,"icdId":17,"date":"4/5/2017"},
{"cptId":9,"claimId":58,"patientId":3,"npId":40062,"cost":426,"ndc":3,"icdId":19,"date":"4/28/2017"},
{"cptId":2,"claimId":59,"patientId":13,"npId":40071,"cost":568,"ndc":5,"icdId":3,"date":"4/21/2017"},
{"cptId":1,"claimId":60,"patientId":19,"npId":40078,"cost":853,"ndc":16,"icdId":4,"date":"4/10/2017"},
{"cptId":5,"claimId":61,"patientId":2,"npId":40114,"cost":459,"ndc":11,"icdId":15,"date":"4/5/2017"},
{"cptId":8,"claimId":62,"patientId":2,"npId":40118,"cost":507,"ndc":18,"icdId":17,"date":"4/1/2017"},
{"cptId":3,"claimId":63,"patientId":22,"npId":40119,"cost":517,"ndc":20,"icdId":15,"date":"4/9/2017"},
{"cptId":10,"claimId":64,"patientId":26,"npId":40134,"cost":407,"ndc":16,"icdId":10,"date":"4/5/2017"},
{"cptId":2,"claimId":65,"patientId":16,"npId":50006,"cost":694,"ndc":18,"icdId":14,"date":"4/22/2017"},
{"cptId":3,"claimId":66,"patientId":26,"npId":50007,"cost":687,"ndc":16,"icdId":12,"date":"4/26/2017"},
{"cptId":9,"claimId":67,"patientId":13,"npId":50013,"cost":838,"ndc":11,"icdId":20,"date":"4/23/2017"},
{"cptId":2,"claimId":68,"patientId":6,"npId":50017,"cost":692,"ndc":17,"icdId":12,"date":"4/16/2017"},
{"cptId":10,"claimId":69,"patientId":20,"npId":50026,"cost":636,"ndc":19,"icdId":19,"date":"4/27/2017"},
{"cptId":1,"claimId":70,"patientId":1,"npId":50036,"cost":621,"ndc":18,"icdId":11,"date":"4/10/2017"},
{"cptId":5,"claimId":71,"patientId":28,"npId":50039,"cost":597,"ndc":19,"icdId":11,"date":"4/23/2017"},
{"cptId":7,"claimId":72,"patientId":16,"npId":50056,"cost":664,"ndc":11,"icdId":2,"date":"4/28/2017"},
{"cptId":3,"claimId":73,"patientId":18,"npId":50057,"cost":625,"ndc":10,"icdId":9,"date":"4/4/2017"},
{"cptId":2,"claimId":74,"patientId":26,"npId":50058,"cost":1006,"ndc":2,"icdId":16,"date":"4/11/2017"},
{"cptId":6,"claimId":75,"patientId":15,"npId":50069,"cost":635,"ndc":18,"icdId":14,"date":"4/17/2017"},
{"cptId":1,"claimId":76,"patientId":21,"npId":50077,"cost":793,"ndc":8,"icdId":20,"date":"4/29/2017"},
{"cptId":7,"claimId":77,"patientId":16,"npId":50082,"cost":639,"ndc":16,"icdId":1,"date":"4/24/2017"},
{"cptId":4,"claimId":78,"patientId":17,"npId":50084,"cost":776,"ndc":6,"icdId":3,"date":"4/8/2017"},
{"cptId":3,"claimId":79,"patientId":14,"npId":50093,"cost":658,"ndc":17,"icdId":4,"date":"4/24/2017"},
{"cptId":3,"claimId":80,"patientId":2,"npId":50107,"cost":697,"ndc":14,"icdId":10,"date":"4/15/2017"},
{"cptId":8,"claimId":81,"patientId":18,"npId":50108,"cost":791,"ndc":20,"icdId":10,"date":"4/17/2017"},
{"cptId":6,"claimId":82,"patientId":25,"npId":50115,"cost":558,"ndc":16,"icdId":12,"date":"4/8/2017"},
{"cptId":8,"claimId":83,"patientId":13,"npId":50128,"cost":660,"ndc":18,"icdId":8,"date":"4/15/2017"},
{"cptId":1,"claimId":84,"patientId":8,"npId":50133,"cost":695,"ndc":5,"icdId":2,"date":"4/22/2017"},
{"cptId":8,"claimId":85,"patientId":8,"npId":50168,"cost":547,"ndc":16,"icdId":19,"date":"4/28/2017"},
{"cptId":1,"claimId":86,"patientId":2,"npId":50174,"cost":777,"ndc":19,"icdId":7,"date":"4/9/2017"},
{"cptId":5,"claimId":87,"patientId":24,"npId":50179,"cost":689,"ndc":17,"icdId":6,"date":"4/8/2017"},
{"cptId":8,"claimId":88,"patientId":12,"npId":50485,"cost":559,"ndc":17,"icdId":19,"date":"4/14/2017"},
{"cptId":10,"claimId":89,"patientId":25,"npId":50180,"cost":678,"ndc":9,"icdId":13,"date":"4/19/2017"},
{"cptId":2,"claimId":90,"patientId":4,"npId":50197,"cost":608,"ndc":18,"icdId":15,"date":"4/19/2017"},
{"cptId":10,"claimId":91,"patientId":1,"npId":50222,"cost":741,"ndc":10,"icdId":19,"date":"4/4/2017"},
{"cptId":1,"claimId":92,"patientId":24,"npId":50224,"cost":539,"ndc":6,"icdId":9,"date":"4/7/2017"},
{"cptId":6,"claimId":93,"patientId":21,"npId":50225,"cost":628,"ndc":14,"icdId":4,"date":"4/19/2017"},
{"cptId":3,"claimId":94,"patientId":12,"npId":50232,"cost":625,"ndc":12,"icdId":1,"date":"4/2/2017"},
{"cptId":5,"claimId":95,"patientId":17,"npId":50235,"cost":764,"ndc":15,"icdId":9,"date":"4/17/2017"},
{"cptId":2,"claimId":96,"patientId":4,"npId":50238,"cost":702,"ndc":15,"icdId":13,"date":"4/10/2017"},
{"cptId":9,"claimId":97,"patientId":30,"npId":50243,"cost":590,"ndc":11,"icdId":2,"date":"4/20/2017"},
{"cptId":4,"claimId":98,"patientId":1,"npId":50262,"cost":1095,"ndc":18,"icdId":7,"date":"4/29/2017"},
{"cptId":1,"claimId":99,"patientId":23,"npId":50280,"cost":725,"ndc":18,"icdId":5,"date":"4/6/2017"},
{"cptId":2,"claimId":100,"patientId":22,"npId":50290,"cost":551,"ndc":2,"icdId":12,"date":"4/24/2017"},


    ]
    self.claims = new oj.ArrayTableDataSource(claims2,{idAttribute:"claimId"});
    
    self.mapDs = [];
     self.visitCostSeries = ko.observableArray([]);
     self.visitCostGroup  = ko.observableArray([]);
    self.showPatientMetrics = function(data,event){
        if (event.target.id !== null){
            var x = []; var y =[];var z=[];var patId='';
            claims2.forEach(function(data){
                  if (data.patientId.toString()  === event.target.id.toString()){
                      patId = data.patientId.toString() + " cost";
                      x.push(data.cptId);
                      y.push(data.cost);
                      z.push(avgCost(claims2,data.cptId,data.patientId));
                  }
            });
            var pvcSeries = [{"name":"Patient cost" ,"items":y},
                             {"name":"National average cost","items":z}
                            ];
            self.visitCostSeries(pvcSeries);
            self.visitCostGroup(x);
            $('#patientPopup').ojPopup('open');
        }
    };
    function avgCost(claims,pId,patId){
        var val = 0; var cnt =0;
        claims.forEach(function(data2){
                if(data2.cptId.toString() === pId.toString() && data2.patientId.toString() !== patId.toString()){
                    val = val+data2.cost;
                    cnt++;
                }
            });
            return val/cnt;
    }
   /* var pvcSeries = [{"name":"visit","items":[100,200,250,430,550]},
                    {"name":"followup","items":[50,250,350,40,150]},
                    ];
    var pvcGroup  = ['visit1','visit2','visit3','visit4','visit5'];*/
    //self.initialize = function(){
       
    //};
    return new DashboardViewModel();
  }
);
