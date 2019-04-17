from flask_wtf import Form
from input_validators import multiple_of
from wtforms import (
    BooleanField,
    DateField,
    FieldList,
    FileField,
    FloatField,
    FormField,
    IntegerField,
    TextField,
)
from wtforms.validators import (
    DataRequired,
    Length,
    NumberRange,
    Regexp,
    required,
)


class Directory(Form):

    name = TextField(validators=[DataRequired(message="")])
