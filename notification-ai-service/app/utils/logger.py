import logging
import sys

# Configure logger
logger = logging.getLogger('notification_ai_service')
logger.setLevel(logging.DEBUG)

# Create console handler
handler = logging.StreamHandler(sys.stdout)
handler.setLevel(logging.DEBUG)

# Create formatter
formatter = logging.Formatter('[NOTIFICATION_AI_SERVICE] %(asctime)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)

# Add handler to logger
if not logger.handlers:
    logger.addHandler(handler)

def info(msg: str):
    logger.info(msg)

def error(msg: str):
    logger.error(msg)

def debug(msg: str):
    logger.debug(msg)

def warning(msg: str):
    logger.warning(msg)