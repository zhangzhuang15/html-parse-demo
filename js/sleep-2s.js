(function(){
    const iframe = document.getElementById('iframe')

    if (iframe) {
        console.log('index.html: sleep 2s, iframe can be accessed')
    } else {
        console.log("index.html: sleep 2s, iframe not found")
    }
})()

