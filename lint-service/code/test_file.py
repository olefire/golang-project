import os
import tkinter as tk
from tkinter import filedialog
from tkinter import ttk

import numpy as np
import pandas as pd

iterations = []


def mac_method(data):
    data = data.apply(lambda x: x - x.min(), axis=1)

    zero_index = np.argwhere(data.to_numpy() == 0)
    unique_index = from_the_beginning(zero_index)

    if len(unique_index) != len(data):
        unique_index = from_the_end(zero_index)

    if len(unique_index) != len(data):
        data = data.apply(lambda x: x - x.min(), axis=0)

        zero_index = np.argwhere(data.to_numpy() == 0)
        unique_index = from_the_beginning(zero_index)

        if len(unique_index) != len(data):
            unique_index = from_the_end(zero_index)

            if len(unique_index) != len(data):
                rows_to_cross_out = np.where(data.apply(lambda x: np.sum(x == 0) > 1, axis=1))[0]
                cols_to_cross_out = \
                    np.where(data.drop(index=rows_to_cross_out).apply(lambda x: np.sum(x == 0) > 0, axis=0))[0]

                min_from_table = data.drop(index=rows_to_cross_out, columns=cols_to_cross_out).min().min()
                data.iloc[np.ix_(rows_to_cross_out, cols_to_cross_out)] += min_from_table
                data.iloc[np.ix_(np.setdiff1d(np.arange(len(data)), rows_to_cross_out),
                                 np.setdiff1d(np.arange(len(data)), cols_to_cross_out))] -= min_from_table

                zero_index = np.argwhere(data.to_numpy() == 0)
                unique_index = from_the_beginning(zero_index)

                if len(unique_index) != len(data):
                    unique_index = from_the_end(zero_index)

                    if len(unique_index) != len(data):
                        return mac_method(data)

    return pd.DataFrame(unique_index, columns=['row', 'col'])


def from_the_beginning(x, i=None, j=None, index=None):
    if i is None:
        i = set()
    if j is None:
        j = set()
    if index is None:
        index = []

    find_zero = x[~np.isin(x[:, 0], list(i)) & ~np.isin(x[:, 1], list(j))]

    if len(find_zero) > 1:
        i.add(find_zero[0, 0])
        j.add(find_zero[0, 1])
        index.append(find_zero[0])

        return from_the_beginning(find_zero, i, j, index)
    else:
        return index + find_zero.tolist()


def from_the_end(x, i=None, j=None, index=None):
    if i is None:
        i = set()
    if j is None:
        j = set()
    if index is None:
        index = []

    find_zero = x[~np.isin(x[:, 0], list(i)) & ~np.isin(x[:, 1], list(j))]

    if len(find_zero) > 1:
        i.add(find_zero[-1, 0])
        j.add(find_zero[-1, 1])
        index.append(find_zero[-1])

        return from_the_end(find_zero, i, j, index)
    else:
        return index + find_zero.tolist()


def browse_file():
    filename = filedialog.askopenfilename(filetypes=[("CSV files", "*.csv")])
    if filename:
        file_path.set(filename)


def perform_calculation():
    table = pd.read_csv(file_path.get(), header=0, index_col=0, sep=";")
    unique_index = mac_method(table)
    optimal_value = sum(table.iloc[row, col] for row, col in zip(unique_index['row'], unique_index['col']))
    result_text = f"\n\nОптимальное значение - {optimal_value}"

    result.set(result_text)


if __name__ == '__main__':
    root = tk.Tk()
    root.title("Метод мака")

    mainframe = ttk.Frame(root, padding="30")
    mainframe.grid(column=0, row=0, sticky='WENS')

    file_path = tk.StringVar()
    optimization = tk.BooleanVar()
    trace = tk.BooleanVar()

    ttk.Label(mainframe, text="Выберите CSV файл: ").grid(column=1, row=1, sticky='WE')
    ttk.Entry(mainframe, textvariable=file_path, width=40).grid(column=2, row=1, sticky='WE')
    ttk.Button(mainframe, text="Обзор", command=browse_file).grid(column=3, row=1, sticky='WE')
    ttk.Button(mainframe, text="Рассчитать", command=perform_calculation).grid(column=2, row=4, sticky='WE')

    result = tk.StringVar()
    ttk.Label(mainframe, text="Результат:").grid(column=1, row=4, sticky='WE')
    ttk.Label(mainframe, textvariable=result, wraplength=250).grid(column=2, row=5, sticky='WE')

    save_status = tk.StringVar()
    ttk.Label(mainframe, textvariable=save_status, wraplength=250).grid(column=2, row=6, sticky='WE')

    root.mainloop()
