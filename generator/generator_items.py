

from generator.models import Item, Item_Category

DEFAULT_ITEMS = [
    Item(1, "Red wine", 12.0, Item_Category.ALCOHOL),
    Item(2, "white wine", 8.0, Item_Category.ALCOHOL),
    Item(3, "Beer", 2.0, Item_Category.ALCOHOL),

    Item(4, "Bread", 1.50, Item_Category.FOOD),
    Item(5, "Candy", 0.5, Item_Category.FOOD),
    

]


def generate_items(number_of_items: int, number_of_categories: int) -> list[Item]:
    '''
    Generates the number of specified items
    '''
