function searchOnClick(event, value) {
  if(event.keyCode == 13) {
    URLString = value.replace(/\s/gm, '+')
    window.location.href = 'http://localhost:1357/translate?' + '-t=' + URLString
  }
}