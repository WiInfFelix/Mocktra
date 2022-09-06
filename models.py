from dataclasses import dataclass

@dataclass
class Item:
    '''Class representing an Item'''
    id: int
    name: str
    value: float


@dataclass
class ItemPurchase:
    '''ItemPurchase represents the purchase of a single item '''
    id: int
    item_id: int
    amount: int


@dataclass
class Transaction:
    '''Transaction bundles together purchased items in single Transactions'''
    id: int
    item_purchases: list[ItemPurchase]


@dataclass
class Location:
    '''Location represents '''
    id: int
    name: str
    transactions: list[Transaction]
