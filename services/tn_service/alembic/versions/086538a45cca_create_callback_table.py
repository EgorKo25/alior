"""create callback table

Revision ID: 086538a45cca
Revises: 
Create Date: 2024-06-23 14:38:33.556769

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '086538a45cca'
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    op.create_table(
        'callbacks',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('name', sa.String(50), nullable=False),
        sa.Column('date', sa.String(50), nullable=False),
        sa.Column('number', sa.String(50), nullable=False),
    )


def downgrade() -> None:
    op.drop_table('callbacks')