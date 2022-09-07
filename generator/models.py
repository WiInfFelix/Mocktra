from dataclasses import dataclass
from enum import Enum


class Item_Category(Enum):
    ALCOHOL = 1
    FOOD = 2
    HYGIENE = 3


@dataclass
class Item:
    '''Class representing an Item'''
    id: int
    name: str
    value: float
    category: Item_Category


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
