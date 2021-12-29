ymaps.ready(init);

// Initialize and add the map
function init() {
    var kazakhstan = [51.169392, 71.449074] ;

    var map = new ymaps.Map('map', {
        zoom: 1,
        center: kazakhstan,
        controls: ['zoomControl']
      });
   
    //Add Marker Function
    function addMarker(coords, name) {
        var placemark = new ymaps.GeoObject(
            {
            geometry: {
                type: "Point",
                coordinates: coords
            },
            properties: 
            {
            hintContent: name,
            balloonContent: name
            }
        })
        map.geoObjects.add(placemark);
    }

    let elements = document.getElementsByClassName("city")
    Array.from(elements).forEach(function (element) {
        var myGeocoder = ymaps.geocode(element.id.replace('-', ', ').replace('_', ', '));
        myGeocoder.then(
            function (res) {
                addMarker(res.geoObjects.get(0).geometry.getCoordinates(), element.id.replace('-', ', ').replaceAll('_', ' '));
            },
        );
    });
  }
  