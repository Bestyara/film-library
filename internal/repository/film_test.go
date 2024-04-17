package repository

import (
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SelectFilm(t *testing.T) {
	t.Parallel()
	t.Run("smoke test", func(t *testing.T) {
		//arrange
		m := mockinit(t)
		defer m.ctrl.Finish()
		m.mockDB.EXPECT().Get(ctx, &film, `SELECT id,name,description,rating,releasedate FROM films WHERE id = $1`, id).Return(nil)
		//act
		f, err := m.repo.SelectFilm(ctx, id)
		//assert
		require.NoError(t, err)
		assert.Equal(t, f.ID, film.ID)

	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			//arrange
			m := mockinit(t)
			defer m.ctrl.Finish()
			m.mockDB.EXPECT().Get(gomock.Any(), gomock.Any(), `SELECT id,name,description,rating,releasedate FROM films WHERE id = $1`, gomock.Any()).Return(sql.ErrNoRows)
			//act
			f, err := m.repo.SelectFilm(ctx, id)
			//assert
			require.EqualError(t, err, "not found")
			require.True(t, errors.Is(err, ErrObjectNotFound))
			assert.Nil(t, f)
		})
		t.Run("internal error", func(t *testing.T) {
			//arrange
			m := mockinit(t)
			defer m.ctrl.Finish()
			m.mockDB.EXPECT().Get(gomock.Any(), gomock.Any(), `SELECT id,name,description,rating,releasedate FROM films WHERE id = $1`, gomock.Any()).Return(assert.AnError)
			//act
			f, err := m.repo.SelectFilm(ctx, id)
			//assert
			require.EqualError(t, err, "assert.AnError general error for testing")
			assert.Nil(t, f)
		})
	})
}
