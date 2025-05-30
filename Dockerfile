# 使用官方輕量版 Python 3.11
FROM python:3.11-alpine

# 設定工作目錄
WORKDIR /app

# 複製依賴檔案
COPY requirements.txt .

# 安裝 Python 需要的套件
RUN pip install --no-cache-dir -r requirements.txt

# 複製應用程式
COPY app.py .



EXPOSE 9090

# 啟動指令
CMD ["python", "app.py"]


