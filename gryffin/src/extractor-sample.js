// Derived from casperjs
function triggerMouseEvent(el, type) {

    try {
        var evt = document.createEvent("MouseEvents"), // HL
            center_x = 1, center_y = 1; 
        try {
            var pos = el.getBoundingClientRect();
            center_x = Math.floor((pos.left + pos.right) / 2);
            center_y = Math.floor((pos.top + pos.bottom) / 2);
        } catch(e) {}
        evt.initMouseEvent( // HL
            type, true, true, window, 1, 1, 1, center_x, center_y, false, false, 
            false, false, 0, el);
        el.dispatchEvent(evt);
        return true;
    } catch (e) {
        return false;
    }
};

function observeDOMChanges(onNewNode) {
    // create an observer instance
    var observer = new window.MutationObserver( // HL
        function(mutations) {
            mutations.forEach(function(mutation) {
                onNewNode && [].forEach.call(mutation.addedNodes || [mutations.target], 
                    function(node){
                        node && (node.nodeType === 1) && onNewNode.call(this, node, observer);

    // ... 

}

jsLinks.getData = function() {
    // self = window.document
    if (self.links.length)
        results.links = self.links;  // HL
    if (self.forms.length)
        results.forms = self.forms;  // HL
    // ...
}