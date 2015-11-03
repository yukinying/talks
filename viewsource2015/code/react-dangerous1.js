    function createMarkup() { return {__html: 'First &middot; Second'}; };
    <div dangerouslySetInnerHTML={createMarkup()} />
