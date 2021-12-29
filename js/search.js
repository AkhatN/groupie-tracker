const searchBar = document.getElementById("search")
                    searchBar.addEventListener("input", (event) => {
                    const info = document.getElementById("info")
                        var inp = searchBar.value.split(' -> ');
                        if (inp[1] == "Locations") {
                          searchBar.value = inp[0]
                          info.selectedIndex = 2
                        };
                        if (inp[1] == "Members") {
                          searchBar.value = inp[0]
                          info.selectedIndex = 1
                        };
                        if (inp[1] == "Artist/Band") {
                          searchBar.value = inp[0]
                          info.selectedIndex = 0
                        };
                        if (inp[1] == "First album date") {
                          searchBar.value = inp[0]
                          info.selectedIndex = 3
                        };
                        if (inp[1] == "Creation date") {
                          searchBar.value = inp[0]
                          info.selectedIndex = 4
                        };
                    });