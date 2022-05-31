const SET_LIST_DADOS_GENERICO = (state, obj) => {
    if (typeof obj.nome === 'object') {
        let key = Object.keys(obj.nome)[0];
        state[key][obj.nome[key]] = obj.lista;
    } else {
        state[obj.nome] = obj.lista;
    }
};

const DELETE_ITEM_FROM_LIST = (state, obj) => {
    let index = -1;
    if (typeof obj.list_name === 'object') { // deleta iten do state do modulo passado em obj
        let key = Object.keys(obj.list_name)[0];
        index = state[key][obj.list_name[key]].findIndex(item => item.id == obj.id);
        if (index > -1) {
            state[key][obj.list_name[key]].splice(index, 1);
        }
    } else { // deleta iten da state global
        index = state[obj.list_name].findIndex(item => item.id == obj.id);
        if (index > -1) {
            state[obj.list_name].splice(index, 1);
        }
    }
};

const SET_GLOBAL_ERROS = (state, obj) => {
    if (state.global_erros.hasOwnProperty(obj.owner)){
        state.global_erros[obj.owner][obj.tipo] = obj.erro;
    } else {
        state.global_erros[obj.owner] = {};
        state.global_erros[obj.owner][obj.tipo] = {};
        state.global_erros[obj.owner][obj.tipo] = obj.erro;
    }
};

export default {
    SET_LIST_DADOS_GENERICO,
    DELETE_ITEM_FROM_LIST,
    SET_GLOBAL_ERROS,
}
