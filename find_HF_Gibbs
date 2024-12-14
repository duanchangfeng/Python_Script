import os
import re
import glob

# 打开文件并遍历行
def find_lines(file_path,txt_path):
    res = {}
    try:
        log_files = glob.glob(os.path.join(file_path, "*.log"))

        if not log_files:
            print("No .log files found in the specified folder.")
            return
        else:
            for log_file in log_files:
                print(f"正在处理文件: {log_file}")
                with open(log_file, 'r', encoding='utf-8') as file:

                    HF = None
                    Gibbs = None

                    for line_number, line in enumerate(file, start=1):

                        # 使用正则表达式寻找 \HF=... 的模式
                        match = re.search(r"\\HF=(-?\d+\.\d+)", line)
                        if match:
                            HF = float(match.group(1))  # 更新为最后一次匹配的数字


                        # 检查是否包含 'Gibbs'
                        if 'Gibbs' in line:
                            # print(f"Line {line_number}: {line.strip()}")
                            try:
                                # 使用正则表达式提取最后的数字
                                match = re.search(r"-?\d+\.\d+$", line.strip())
                                if match:
                                    Gibbs = float(match.group())  # 返回匹配的数字并转换为浮点数
                                else:
                                    print("No number found at the end of the line.")
                                    return None
                            except Exception as e:
                                print(f"An error occurred: {e}")
                                return None

                    res[os.path.basename(log_file)] = [HF,Gibbs]


                    print("HF",HF)
                    print("Gibbs",Gibbs)
        print(res)
        # 文件输出
        with open(txt_path, "w", encoding="utf-8") as file:
            for key, value in res.items():
                file.write(f"{str(key)} {str(value[0]).rjust(20)} {str(value[1]).rjust(20)}\n")
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")

# 替换为你的输入文件夹路径
# log_folder_path = 'C:\\Users\\PC\\Desktop\\test'

log_folder_path = os.getcwd()
# 替换为你的输出文件
txt_folder_path = log_folder_path + '\\output.txt'

find_lines(log_folder_path,txt_folder_path)


