# 1p-notes
A note taking app with a 1password backend.

Ok, let's see..

So it's gonna be some kind of GUI that reads-, and writes notes to a local 1password instance. Let's list some requirements:

- global fuzzy search function
- tabbed notes view
- keyboard shortcuts
- read, save and update notes

The GUI could either be some web app or a gnome gedit plugin.

````sh
# http api
GET    /note
GET    /note?title=<title>
POST   /note w json body
````

