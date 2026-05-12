import logging
import io
import random
from telegram import Update
from telegram.ext import (
    ApplicationBuilder,
    CommandHandler,
    MessageHandler,
    filters,
    ContextTypes,
)
from docx import Document

logging.basicConfig(
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
    level=logging.INFO,
)
logger = logging.getLogger(__name__)

BOT_TOKEN = "8770712457:AAEis1XxiOFywgpjMWwrpAKTJ3lN9qT3hh8"


def parse_questions(docx_bytes: bytes) -> list[dict]:
    """
    Jadval strukturasi:
    - 1-ustun: raqam (1,2,3...) yoki harf (A,B,C,D)
    - 2-ustun: savol matni yoki variant matni
    - To'g'ri javob 2-ustunda '#' bilan boshlanadi
    """
    doc = Document(io.BytesIO(docx_bytes))

    questions = []

    for table in doc.tables:
        current_question = None
        current_options = []
        correct_answer = None

        for row in table.rows:
            if len(row.cells) < 2:
                continue

            key = row.cells[0].text.strip()
            value = row.cells[1].text.strip()

            if not key or not value:
                continue

            # Savol qatori: raqam (1, 2, 3 ...)
            if key.isdigit():
                # Oldingi savolni saqlash
                if current_question and len(current_options) == 4 and correct_answer:
                    options = current_options[:]
                    random.shuffle(options)
                    correct_index = options.index(correct_answer)
                    questions.append({
                        "question": current_question,
                        "options": options,
                        "correct_option_id": correct_index,
                    })

                # Yangi savol boshlash
                current_question = value
                current_options = []
                correct_answer = None

            # Variant qatori: A, B, C, D
            elif key.upper() in ("A", "B", "C", "D"):
                if value.startswith("#"):
                    clean = value[1:].strip()
                    correct_answer = clean
                    current_options.append(clean)
                else:
                    current_options.append(value)

        # Oxirgi savolni saqlash
        if current_question and len(current_options) == 4 and correct_answer:
            options = current_options[:]
            random.shuffle(options)
            correct_index = options.index(correct_answer)
            questions.append({
                "question": current_question,
                "options": options,
                "correct_option_id": correct_index,
            })

    return questions


async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text(
        "👋 Salom! Men Quiz Bot!\n\n"
        "📄 Menga <b>.docx</b> fayl yuboring.\n\n"
        "Fayl jadval ko'rinishida bo'lsin:\n"
        "• 1-ustun: raqam yoki variant harfi (A/B/C/D)\n"
        "• 2-ustun: savol yoki variant matni\n"
        "• To'g'ri javob <b>#</b> belgisi bilan boshlansin\n\n"
        "Bot variantlarni <b>aralashtiradi</b> va Quiz Poll yuboradi! 🎯",
        parse_mode="HTML",
    )


async def handle_document(update: Update, context: ContextTypes.DEFAULT_TYPE):
    doc = update.message.document

    if not doc.file_name.endswith(".docx"):
        await update.message.reply_text(
            "❌ Noto'g'ri fayl formati!\n"
            "Iltimos, faqat <b>.docx</b> fayl yuboring.",
            parse_mode="HTML",
        )
        return

    await update.message.reply_text("⏳ Fayl o'qilmoqda, savollar tayyorlanmoqda...")

    try:
        file = await context.bot.get_file(doc.file_id)
        file_bytes = await file.download_as_bytearray()
    except Exception as e:
        logger.error(f"Fayl yuklab olishda xato: {e}")
        await update.message.reply_text("❌ Faylni yuklab olishda xatolik yuz berdi.")
        return

    try:
        questions = parse_questions(bytes(file_bytes))
    except Exception as e:
        logger.error(f"Faylni o'qishda xato: {e}")
        await update.message.reply_text(
            "❌ Faylni o'qishda xatolik!\n"
            "Fayl to'g'ri formatda ekanligini tekshiring."
        )
        return

    if not questions:
        await update.message.reply_text(
            "⚠️ Faylda hech qanday savol topilmadi!\n"
            "To'g'ri javob oldiga <b>#</b> belgisi qo'yilganini tekshiring.",
            parse_mode="HTML",
        )
        return

    await update.message.reply_text(
        f"✅ <b>{len(questions)} ta savol</b> topildi!\nQuiz boshlanmoqda... 🎯",
        parse_mode="HTML",
    )

    sent = 0
    for i, q in enumerate(questions):
        try:
            await context.bot.send_poll(
                chat_id=update.effective_chat.id,
                question=f"{i + 1}. {q['question']}",
                options=q["options"],
                type="quiz",
                correct_option_id=q["correct_option_id"],
                is_anonymous=False,
            )
            sent += 1
        except Exception as e:
            logger.error(f"{i+1}-savol yuborishda xato: {e}")
            await update.message.reply_text(f"⚠️ {i+1}-savol yuborishda xatolik: {e}")

    await update.message.reply_text(
        f"🎉 Quiz tugadi!\n✅ {sent} ta savol yuborildi."
    )


async def handle_other(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text(
        "📄 Iltimos, menga faqat <b>.docx</b> fayl yuboring.\n"
        "/start — yordam olish uchun",
        parse_mode="HTML",
    )


def main():
    app = ApplicationBuilder().token(BOT_TOKEN).build()

    app.add_handler(CommandHandler("start", start))
    app.add_handler(MessageHandler(filters.Document.ALL, handle_document))
    app.add_handler(MessageHandler(~filters.Document.ALL & ~filters.COMMAND, handle_other))

    logger.info("Bot ishga tushdi...")
    app.run_polling()


if __name__ == "__main__":
    main()
