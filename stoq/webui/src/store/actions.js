import config from '../config';


const setListPaisesAction = async (store, obj) => {
    if (store.state.lista_paises.length <= 0) {
        let url = config.localidade + 'pais';
        try {
            const { data } = await Http.getData(url);
            store.commit('SET_LIST_DADOS_GENERICO', {
                nome: 'lista_paises',
                lista: data['paises']
            });
        } catch(e) {
            console.log('setListPaisesAction', e);
        }
    }
};

const setListEstadosAction = async (store, obj) => {
    if (store.state.lista_estados.length <= 0) {
        let url = config.localidade + 'estado';
        try {
            const { data } = await Http.getData(url);
            store.commit('SET_LIST_DADOS_GENERICO', {
                nome: 'lista_estados',
                lista: data['estados']
            });
        } catch(e) {
            console.log('setListEstadosAction', e);
        }
    }
};

const setListMunicipiosUFAction = async (store, obj) => {
    if (!store.state.lista_municipios_geral.hasOwnProperty(obj.uf)) {
        let url = config.localidade + 'municipio?uf=' + obj.uf;
        try {
            const { data } = await Http.getData(url);
            store.commit('SET_LIST_DADOS_GENERICO', {
                nome: {lista_municipios_geral: obj.uf},
                lista: data['municipios']
            });
            store.commit('SET_LIST_DADOS_GENERICO', {
                nome: 'lista_municipios',
                lista: data['municipios']
            });
        } catch(e) {
            console.log('setListMunicipiosUFAction', e);
        }
    } else {
        let dados = store.state.lista_municipios_geral[obj.uf];
        store.commit('SET_LIST_DADOS_GENERICO', {
            nome: 'lista_municipios',
            lista: dados
        });
    }
};

const deleteItemFromListAction = async (store, obj) => {
    store.commit('DELETE_ITEM_FROM_LIST', {
        id: obj.id,
        list_name: obj.list_name,
    });
};


export default {
    setListPaisesAction,
    setListEstadosAction,
    setListMunicipiosUFAction,
    deleteItemFromListAction,
};
