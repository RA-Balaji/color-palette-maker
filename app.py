from flask import Flask, render_template, request, jsonify
from colorthief import ColorThief
app = Flask(__name__)

def get_color_palette(img, num):
    color_thief = ColorThief(img)
    palette = color_thief.get_palette(color_count=num)
    palette_hex = []
    for rgb in palette:
        palette_hex.append(rgb_to_hex(rgb))
    return palette_hex
def rgb_to_hex(rgb):
    return '#{:02x}{:02x}{:02x}'.format(rgb[0], rgb[1], rgb[2])

@app.route('/', methods=['GET', 'POST'])
def upload_file():
    if request.method == 'POST':
        file = request.files['file']
        num_colors = int(request.form['num_colors'])
        palette = get_color_palette(file, num_colors+1)
        return jsonify({'palette': palette})
    return render_template('index.html')

if __name__ == '__main__':
    app.run(debug=True)
