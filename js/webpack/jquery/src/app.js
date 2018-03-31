$(document).ready(() => {
    $('#testText').text('hello')

    let count = 0
    $('#testButton').click(_.throttle(() => {
        count++
        $('#testText').text('push:' + count)
    }, 1000))
})