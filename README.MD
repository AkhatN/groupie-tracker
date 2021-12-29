# Groupie tracker


## Description

Groupie Trackers receives a given API and manipulates the data contained in it, and displays the information in the site.

It is given an API, that consists in two parts:

The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

And the second one, relation, does the link between all the other parts, artists, dates and locations.

It is a user friendly website where you can see the bands info through several data visualizations (examples : blocks, cards, pages). 


Groupie tracker has some additional features:
* search bar
* filters
* geolocalization


Search bar searches inside the website for a specific text input. A client can search a member or artist or any other attribute in the data system.

The program handles these search cases :
    * artist/band name
    * members
    * locations
    * first album date
    * creation date

The search bar has typing suggestions as you write.
    The search bar identifies and displays in each suggestion the individual type of the search cases. (ex: Freddie Mercury -> member)
    For example if you start writing "phil" it appears as suggestions Phil Collins - member and Phil Collins - artist/band. This is just an example of a display.



Filters allows the user filter the artists/bands that will be shown.

The project incorporates these four filters:
    * filter by creation date
    * filter by first album date
    * filter by number of members
    * filter by locations of concerts




Geolocalization allows mapping the different concerts locations of a certain artist/band given by the Client.
    Yandex Map Api is used.

   The project can convert addresses (ex: Germany Mainz) into geographic coordinates (ex: 49,59380 8,15052) which places markers for the concerts locations of a certain artist/band on a map.

## How to run 
1. Type **go run main.go**
2. Open **localhost:8080** in any browser