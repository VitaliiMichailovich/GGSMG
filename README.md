# GGSMG = Golang Gin Site Map Generator
Web site wich goes to every link in your site and give you back a sitemap file on e-mail.
## Build and deploy
After every commit in `main` brunch this code will build with [CircleCI](https://circleci.com/).

If build ok and tests pass it goes to [Heroku](https://heroku.com/) and deploy in [Docker](https://www.docker.com/) container.
## Link 
To look how it works you can go to [http://ggsmg.herokuapp.com/](http://ggsmg.herokuapp.com/) and check it.
## ToDo:
1. Implement giving back an error message from post request. 
1. ~~Implement correct error messages from checkIn package.~~ **DONE 07.11.2017**
1. Create sitemap.xml file generator.
1. Create smtp sender with attaching sitemap.xml file.
1. Tests
1. TBD...
