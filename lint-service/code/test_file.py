#/bin/python
"""This is a module docstring
"""

class SomeClass:
    """This is a class docstring
    """

    @staticmethod
    def some_method():
        """This is a method docstring
        """
        print('hello world')

    def public_metho_one(self):
        """This is a method docstring
        """
        print(self.some_method())