from datetime import datetime

from flask_wtf import Form
from input_validators import multiple_of
from libraries.files.Directory import Directory
from wtforms import (BooleanField, DateField, FieldList, FileField, FloatField,
                     FormField, IntegerField, TextField)
from wtforms.validators import (DataRequired, Length, NumberRange, Regexp,
                                required)


class Place(Form):

    created = FormField(datetime)
    dir = FormField(Directory)
    name = TextField(validators=[DataRequired(message="")])
