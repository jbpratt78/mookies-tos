import SwiftUI

struct ContentView: View {
    var body: some View {
        Menu()
    }
}

struct Menu: View {

    @State private var selection: Set<Tospb_Category> = []
    
    @State var itemSelected: Tospb_Item?

    var menu: Tospb_Menu
    init() {
        self.menu = loadMenu()
    }

    var body: some View {
        ZStack {
            ScrollView {
                VStack(alignment: .leading){
                    ForEach(menu.categories, id: \.self) { cat in
                        CategoryView(category: cat, isExpanded: self.selection.contains(cat), itemSelected: self.$itemSelected)
                            .onTapGesture { self.selectDeselect(cat) }
                            .animation(.linear(duration: 0.2))
                    }.padding()
                }
            }
            
            if self.itemSelected != nil {
                PopupMenu(item: self.itemSelected!)
                    .padding(.horizontal)
            }
        }
    }

    func selectDeselect(_ category: Tospb_Category) {
        if selection.contains(category) {
            selection.remove(category)
        } else {
            selection.insert(category)
        }
    }
}

struct CategoryView: View {
    var category: Tospb_Category

    let isExpanded: Bool
    
    @Binding var itemSelected: Tospb_Item?

    var body: some View {
        HStack {
            content
            Spacer()
        }
        .contentShape(Rectangle())
    }

    private var content: some View {
        VStack(alignment: .leading) {
            Text(category.name).font(.headline)
            if isExpanded {
                ForEach(category.items, id: \.self) { item in
                    HStack {
                        Text(item.name)
                        Spacer()
                        Text(String(format: "%.2f", item.price / 100))
                    }
                    .padding(.top, 10)
                    .onTapGesture {
                        self.itemSelected = item
                        print(item.name)
                    }
                }
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}

struct PopupMenu: View {
    
    var item: Tospb_Item
    
    var body: some View {
        VStack {
            ForEach(item.options, id: \.self) { item in
                OptionView(opt: item, isSelected: true)
                .padding(.top, 10)
            }
            HStack {
                Button(action: {}) {
                    Text("Add to order")
                }
                .padding(.top, 10)
                Spacer()
                Button(action: {}) {
                    Text("Close")
                }
                .padding(.top, 10)
            }
            .buttonStyle(BorderlessButtonStyle())
        }// VStack
        .padding(15)
        .background(RoundedRectangle(cornerRadius: 4)
                        .stroke(Color.black, lineWidth: 2)
                        .background(Color.white))
    }
}

struct OptionView: View {

    var opt: Tospb_Option
    @State var isSelected: Bool

    var body: some View {
        HStack {
            Text(opt.name)
            Spacer()
            Text(String(format: "%.2f", opt.price / 100))
            if isSelected {
                Image(systemName: "checkmark")
            }
        }.onTapGesture {
            isSelected = !isSelected
        }
    }
}

func loadMenu() -> Tospb_Menu {
    guard let menu = try? Tospb_Menu(jsonUTF8Data: Data(jsonMenu.utf8)) else {
        fatalError("failed to decode menu")
    }
    return menu
}

let jsonMenu = """
{
 "categories": [
  {
   "name": "Sandwiches",
   "id": 1,
   "items": [
    {
     "name": "LG Smoked Pulled Pork",
     "price": 495,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "RG Smoked Pulled Pork",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Smoked Chicken Breast",
     "price": 495,
     "options": [
      {
       "name": "pickles"
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo",
       "selected": true
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25,
       "selected": true
      },
      {
       "name": "lettuce",
       "price": 25,
       "selected": true
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "RG Smoked Chicken Breast",
     "price": 395,
     "options": [
      {
       "name": "pickles"
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup",
       "selected": true
      },
      {
       "name": "mayo",
       "selected": true
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25,
       "selected": true
      },
      {
       "name": "lettuce",
       "price": 25,
       "selected": true
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "'The Molly'",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce",
       "selected": true
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25,
       "selected": true
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Hamburger",
     "price": 495,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Hamburger",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Cheeseburger",
     "price": 550,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Cheeseburger",
     "price": 425,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Grilled Cheese",
     "price": 300,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Pulled Pork Melt",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    }
   ]
  },
  {
   "name": "Plates",
   "id": 2,
   "items": [
    {
     "name": "Smoked Pulled Pork",
     "price": 990,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "RG Rib",
     "price": 995,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "½ Smoked Chicken",
     "price": 995,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Smoked Chicken Breast",
     "price": 725,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Smoked Wing (8 wings)",
     "price": 999,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "BBQ Loaded Nachos",
     "price": 875,
     "options": [
      {
       "name": "cheddar cheese",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "jalapeños",
       "selected": true
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      }
     ]
    },
    {
     "name": "Chicken Loaded Nachos",
     "price": 875,
     "options": [
      {
       "name": "cheddar cheese",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "jalapeños",
       "selected": true
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      }
     ]
    }
   ]
  },
  {
   "name": "Baskets",
   "id": 3,
   "items": [
    {
     "name": "Smoked Wing",
     "price": 500,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Rib",
     "price": 500,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    }
   ]
  },
  {
   "name": "Potatoes",
   "id": 4,
   "items": [
    {
     "name": "Loaded Pork",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "Loaded Chicken",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "Loaded Potato (no meat)",
     "price": 699,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "(Pork) Smothered and Covered Fries",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "(Chicken) Smothered and Covered Fries",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    }
   ]
  },
  {
   "name": "Sides",
   "id": 5,
   "items": [
    {
     "name": "SM Cole Slaw",
     "price": 300
    },
    {
     "name": "LG Cole Slaw",
     "price": 600
    },
    {
     "name": "SM Baked Beans",
     "price": 300
    },
    {
     "name": "LG Baked Beans",
     "price": 600
    },
    {
     "name": "SM Potato Salad",
     "price": 300
    },
    {
     "name": "LG Potato Salad",
     "price": 600
    },
    {
     "name": "Plain Chips",
     "price": 100
    },
    {
     "name": "Fries",
     "price": 300
    }
   ]
  },
  {
   "name": "Drinks",
   "id": 6,
   "items": [
    {
     "name": "Canned Drink",
     "price": 100
    },
    {
     "name": "Bottled Water",
     "price": 150
    }
   ]
  },
  {
   "name": "Desserts",
   "id": 7,
   "items": [
    {
     "name": "Oreo Dream",
     "price": 350
    },
    {
     "name": "Lemon Delight",
     "price": 350
    },
    {
     "name": "Strawberry Pizza",
     "price": 350
    },
    {
     "name": "Whole Dessert",
     "price": 3000
    },
    {
     "name": "SM Banana Pudding",
     "price": 350
    },
    {
     "name": "LG Banana Pudding",
     "price": 700
    }
   ]
  },
  {
   "name": "Sauces",
   "id": 8,
   "items": [
    {
     "name": "Extra Ranch",
     "price": 50
    },
    {
     "name": "Extra White",
     "price": 50
    },
    {
     "name": "Extra Buffalo Wing",
     "price": 50
    },
    {
     "name": "Extra BBQ",
     "price": 50
    }
   ]
  }
 ]
}
"""