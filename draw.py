import matplotlib.pyplot as plt
import numpy as np

# figure1


# figure2
def draw_wave(ax, x_start, x_end, y, amplitude=0.3, wavelength=1, points=100):
    """
    绘制波浪线
    :param ax: matplotlib 的绘图对象
    :param x_start: 波浪线的起始 x 坐标
    :param x_end: 波浪线的结束 x 坐标
    :param y: 波浪线的 y 坐标
    :param amplitude: 波浪线的振幅
    :param wavelength: 波长
    :param points: 波浪线的点数量
    """
    x = np.linspace(x_start, x_end, points)
    y_wave = y + amplitude * np.sin(2 * np.pi * (x - x_start) / wavelength)
    ax.plot(x, y_wave, color="blue", lw=1.5)


# 创建绘图
fig, ax = plt.subplots(figsize=(8, 8))

# 节点坐标 (手动指定)
nodes = {
    "a": (0, 6),
    "b": (1, 6),
    "c": (2, 6),
    "d": (1, 4),
    "k": (1, 3),
    "e": (0, 2),
    "f": (2, 2),
    "g": (6, 6),
    "h": (6, 4),
    "i": (6, 3),
    "j": (6, 2),
}

# 绘制连接线
arrow_props = dict(arrowstyle="-", color="darkgray", lw=1.5)
ax.annotate("", xy=(0, 6), xytext=(1, 6), arrowprops=arrow_props)
ax.annotate("", xy=(1, 6), xytext=(2, 6), arrowprops=arrow_props)
ax.annotate("", xy=(1, 3), xytext=(1, 4), arrowprops=arrow_props)
ax.annotate("", xy=(0, 2), xytext=(1, 3), arrowprops=arrow_props)
ax.annotate("", xy=(1, 3), xytext=(2, 2), arrowprops=arrow_props)


arrow_props = dict(arrowstyle="->", color="green", lw=3.5)
ax.annotate("", xy=(4.5, 6), xytext=(3, 6), arrowprops=arrow_props)
ax.annotate("", xy=(4.5, 3), xytext=(3, 3), arrowprops=arrow_props)

# 绘制节点
for label, (x, y) in nodes.items():
    ax.scatter(x, y, color="lightgray", s=500, zorder=3)

# 添加相同的标注到两个节点上
ax.text(0, 6, "l", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(1, 6, "i", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(2, 6, "r", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(1, 4, "a", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(1, 3, "j", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(0, 2, "b", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(2, 2, "c", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")

ax.text(6, 6, "i", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 4, "j", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 3, "j", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 2, "j", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")

ax.text(6, 5.5, r"$x_{lir}$", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 3.5, r"$x_{ajb}$", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 2.5, r"$x_{ajc}$", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")
ax.text(6, 1.5, r"$x_{bjc}$", color="black", fontsize=14, ha="center", va="center",fontname="Georgia")


# 绘制波浪线
draw_wave(ax, 5, 6, 6)
draw_wave(ax, 5, 6, 4)
draw_wave(ax, 5, 6, 3)
draw_wave(ax, 5, 6, 2)

draw_wave(ax, 6, 7, 6)
draw_wave(ax, 6, 7, 4)
draw_wave(ax, 6, 7, 3)
draw_wave(ax, 6, 7, 2)


# 设置图形样式
ax.set_xlim(-1, 8)
ax.set_ylim(-1, 8)
ax.axis("on")
plt.show()




# figure3
import matplotlib.pyplot as plt
import numpy as np


def draw_wave(ax, x_start, x_end, y, amplitude=0.3, wavelength=1, points=100):
    """
    绘制波浪线
    :param ax: matplotlib 的绘图对象
    :param x_start: 波浪线的起始 x 坐标
    :param x_end: 波浪线的结束 x 坐标
    :param y: 波浪线的 y 坐标
    :param amplitude: 波浪线的振幅
    :param wavelength: 波长
    :param points: 波浪线的点数量
    """
    x = np.linspace(x_start, x_end, points)
    y_wave = y + amplitude * np.sin(2 * np.pi * (x - x_start) / wavelength)
    ax.plot(x, y_wave, color="blue", lw=1.5)


# 创建绘图
fig, ax = plt.subplots(figsize=(8, 8))

# 节点坐标 (手动指定)
nodes = {
    "a": (1, 5),

    "c": (1, 3),
    "d": (1, 1.5),
    "k": (1, 0.5),
    "e": (2, 4),
    "f": (4, 4),
    "g": (5, 5),

    "i": (5, 3),
    "j": (5, 1.5),
    'l': (5, 0.5)
}

# 绘制连接线
arrow_props = dict(arrowstyle="-", color="darkgray", lw=1.5)
ax.annotate("", xy=(1, 5), xytext=(2, 4), arrowprops=arrow_props)
ax.annotate("", xy=(1, 3), xytext=(2, 4), arrowprops=arrow_props)
ax.annotate("", xy=(2, 4), xytext=(4, 4), arrowprops=arrow_props)
ax.annotate("", xy=(4, 4), xytext=(5, 5), arrowprops=arrow_props)
ax.annotate("", xy=(4, 4), xytext=(5, 3), arrowprops=arrow_props)

arrow_props = dict(arrowstyle="->", color="green", lw=3.5)
ax.annotate("", xy=(3, 2), xytext=(3, 3), arrowprops=arrow_props)

# 绘制节点
for label, (x, y) in nodes.items():
    ax.scatter(x, y, color="lightgray", s=500, zorder=3)

# 添加相同的标注到两个节点上
ax.text(1, 5, "a", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 3, "b", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 1.5, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 0.5, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(2, 4, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(4, 4, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 5, "c", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 3, "d", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 1.5, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 0.5, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

ax.text(1, 1, r"$x_{aij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 0, r"$x_{bij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 1, r"$x_{ijc}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(5, 0, r"$x_{ijd}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

# 绘制波浪线
draw_wave(ax, 0, 1, 1.5)
draw_wave(ax, 1, 2, 1.5)
draw_wave(ax, 4, 5, 1.5)
draw_wave(ax, 5, 6, 1.5)

draw_wave(ax, 0, 1, 0.5)
draw_wave(ax, 1, 2, 0.5)
draw_wave(ax, 4, 5, 0.5)
draw_wave(ax, 5, 6, 0.5)

# 设置图形样式
ax.set_xlim(-1, 8)
ax.set_ylim(-1, 8)
ax.axis("on")
plt.show()

# figure4
import matplotlib.pyplot as plt
import numpy as np


def draw_wave(ax, x_start, x_end, y, amplitude=0.3, wavelength=1, points=100):
    """
    绘制波浪线
    :param ax: matplotlib 的绘图对象
    :param x_start: 波浪线的起始 x 坐标
    :param x_end: 波浪线的结束 x 坐标
    :param y: 波浪线的 y 坐标
    :param amplitude: 波浪线的振幅
    :param wavelength: 波长
    :param points: 波浪线的点数量
    """
    x = np.linspace(x_start, x_end, points)
    y_wave = y + amplitude * np.sin(2 * np.pi * (x - x_start) / wavelength)
    ax.plot(x, y_wave, color="blue", lw=1.5)


# 创建绘图
fig, ax = plt.subplots(figsize=(8, 8))

# 节点坐标 (手动指定)
nodes = {
    "a": (1, 4),
    "c": (1, 3),
    "d": (1, 2),
    "k": (1, 1),
    "e": (6, 4),
    "f": (6, 3),
    "g": (6, 2),
    "i": (6, 1),
}

arrow_props = dict(arrowstyle="<->", color="green", lw=3.5)
ax.annotate("", xy=(2.5, 4), xytext=(4.5, 4), arrowprops=arrow_props)
ax.annotate("", xy=(2.5, 3), xytext=(4.5, 3), arrowprops=arrow_props)
ax.annotate("", xy=(2.5, 2), xytext=(4.5, 2), arrowprops=arrow_props)
ax.annotate("", xy=(2.5, 1), xytext=(4.5, 1), arrowprops=arrow_props)


# 绘制节点
for label, (x, y) in nodes.items():
    ax.scatter(x, y, color="lightgray", s=500, zorder=3)

# 添加相同的标注到两个节点上
ax.text(1, 4, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 3, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 2, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 1, "i", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

ax.text(6, 4, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 3, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 2, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 1, "j", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")


ax.text(1, 3.5, r"$x_{aij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 2.5, r"$x_{bij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 1.5, r"$x_{aij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(1, 0.5, r"$x_{bij}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

ax.text(6, 3.5, r"$x_{ijc}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 2.5, r"$x_{ijc}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 1.5, r"$x_{ijd}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(6, 0.5, r"$x_{ijd}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

ax.text(3.5, 4.3, r"$\overline{x_{aij}^j,x_{ijc}^i}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(3.5, 3.3, r"$\overline{x_{bij}^j,x_{ijc}^i}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(3.5, 2.3, r"$\overline{x_{aij}^j,x_{ijd}^i}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")
ax.text(3.5, 1.3, r"$\overline{x_{bij}^j,x_{ijd}^i}$", color="black", fontsize=14, ha="center", va="center", fontname="Georgia")

# 绘制波浪线
draw_wave(ax, 0, 1, 4)
draw_wave(ax, 0, 1, 3)
draw_wave(ax, 0, 1, 2)
draw_wave(ax, 0, 1, 1)
draw_wave(ax, 1, 2, 4)
draw_wave(ax, 1, 2, 3)
draw_wave(ax, 1, 2, 2)
draw_wave(ax, 1, 2, 1)

draw_wave(ax, 5, 6, 4)
draw_wave(ax, 5, 6, 3)
draw_wave(ax, 5, 6, 2)
draw_wave(ax, 5, 6, 1)
draw_wave(ax, 6,7, 4)
draw_wave(ax, 6,7, 3)
draw_wave(ax, 6,7, 2)
draw_wave(ax, 6,7, 1)

# 设置图形样式
ax.set_xlim(-1, 8)
ax.set_ylim(-1, 8)
ax.axis("on")
plt.show()
