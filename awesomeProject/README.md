# Query example
 query {
   books {
     title
     author
     published
   }
   students {
     id
     firstname
     lastname
     age
   }
 }

 mutation {
   createBook(title: "My Book", author: "John Doe", published: "2022-01-01") {
     id
     title
     author
     published
   }
 }

query {
studentWithBook {
books {
id
title
author
published
}
students {
id
firstname
lastname
age
}
}
}
